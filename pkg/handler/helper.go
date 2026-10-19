package handler

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// handlerHelperCtx sends a request to the M-Pesa API and decodes the response.
//
// On failure it returns a non-nil *APIError. For backward compatibility the
// response is also returned with its error fields populated, so callers that
// inspect ErrorCode keep working; new code should check err first.
func handlerHelperCtx[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](cl *PokeaClient, ctx context.Context, payload T, url, method, authToken string) (*PokeaResParams, error) {
    const op = "pokea request"

	buff, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

    req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

    var client *http.Client
    if cl != nil {
        client = cl.httpClient
    }
    if client == nil {
        client = &http.Client{}
    }
    res, err := client.Do(req)
    if err != nil {
        return failedResponse(unreachableError(op, 0, nil, err))
    }
    defer res.Body.Close()

    buff, err = io.ReadAll(res.Body)
	if err != nil {
		return failedResponse(unreachableError(op, res.StatusCode, nil, err))
	}

    var result PokeaResParams
    if err := json.Unmarshal(buff, &result); err != nil {
        // Non-JSON body (e.g., HTML error page from a gateway)
        return failedResponse(unreachableError(op, res.StatusCode, buff, err))
    }

    // If key fields are empty, try the B2C callback/result envelope shape
    if result.ResponseCode == "" && result.ErrorCode == "" {
        type b2cResultEnvelope struct {
            Result struct {
                ResultType                 int    `json:"ResultType"`
                ResultCode                 int    `json:"ResultCode"`
                ResultDesc                 string `json:"ResultDesc"`
                OriginatorConversationID   string `json:"OriginatorConversationID"`
                ConversationID             string `json:"ConversationID"`
            } `json:"Result"`
        }
        var env b2cResultEnvelope
        if err := json.Unmarshal(buff, &env); err == nil && (env.Result.ResultDesc != "" || env.Result.ResultCode != 0) {
            // Map to our common response shape per docs
            result.ResponseCode = fmt.Sprintf("%d", env.Result.ResultCode)
            // Prefer the documented ack phrase when unknown
            if env.Result.ResultDesc != "" {
                result.ResponseDescription = env.Result.ResultDesc
            } else if result.ResponseDescription == "" {
                result.ResponseDescription = "Accept the service request successfully."
            }
            result.OriginatorConversationID = env.Result.OriginatorConversationID
            result.ConversationID = env.Result.ConversationID
        }
    }

    if result.ErrorCode != "" || res.StatusCode < 200 || res.StatusCode > 299 {
        apiErr := &APIError{
            Status:    res.StatusCode,
            Code:      result.ErrorCode,
            Message:   result.ErrorMessage,
            RequestID: result.RequestID,
            RawBody:   buff,
            Op:        op,
        }
        if apiErr.Message == "" {
            apiErr.Message = http.StatusText(res.StatusCode)
        }
        return &result, apiErr
    }

    return &result, nil
}

// failedResponse returns a response mirroring apiErr alongside the error itself.
func failedResponse(apiErr *APIError) (*PokeaResParams, error) {
    return &PokeaResParams{
        PokeaErrorParams: PokeaErrorParams{
            RequestID:    apiErr.RequestID,
            ErrorCode:    apiErr.Code,
            ErrorMessage: apiErr.Message,
        },
    }, apiErr
}

// Backward-compatible helper without context
func handlerHelper[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](payload T, url, method, authToken string) (*PokeaResParams, error) {
    cl := &PokeaClient{}
    return handlerHelperCtx(cl, context.Background(), payload, url, method, authToken)
}
