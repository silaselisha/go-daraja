package handler

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func handlerHelperCtx[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](cl *DarajaClient, ctx context.Context, payload T, url, method, authToken string) (*DarajaResParams, error) {
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
        // Backward-compatible: encode network error in response
        return &DarajaResParams{
            DarajaErrorParams: DarajaErrorParams{
                ErrorCode:    "500.003.1001",
                ErrorMessage: "Service is currently unreachable. Please try again later.",
            },
        }, nil
    }

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatalf("failed to close response body %v\n", err)
		}
	}()

    buff, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

    var result DarajaResParams
    if err := json.Unmarshal(buff, &result); err == nil {
        // If we successfully parsed but key fields are empty, try alternate shapes
        if result.ResponseCode == "" && result.ErrorCode == "" {
            // Attempt to parse B2C callback/result envelope
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
                if env.Result.ResultCode == 0 {
                    result.ResponseCode = "0"
                } else {
                    result.ResponseCode = fmt.Sprintf("%d", env.Result.ResultCode)
                }
                // Prefer the documented ack phrase when unknown
                if env.Result.ResultDesc != "" {
                    result.ResponseDescription = env.Result.ResultDesc
                } else if result.ResponseDescription == "" {
                    result.ResponseDescription = "Accept the service request successfully."
                }
                result.OriginatorConversationID = env.Result.OriginatorConversationID
                result.ConversationID = env.Result.ConversationID
                return &result, nil
            }
        } else {
            return &result, nil
        }
    }

    // If the standard response shape failed to parse, try parsing known error fields
    var errPayload DarajaErrorParams
    if err := json.Unmarshal(buff, &errPayload); err == nil {
        result.DarajaErrorParams = errPayload
        return &result, nil
    }

    // Fallback: non-JSON (e.g., HTML error page). Map to a generic unreachable error
    return &DarajaResParams{DarajaErrorParams: DarajaErrorParams{ErrorCode: "500.003.1001", ErrorMessage: "Service is currently unreachable. Please try again later."}}, nil
}

// Backward-compatible helper without context
func handlerHelper[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](payload T, url, method, authToken string) (*DarajaResParams, error) {
    cl := &DarajaClient{}
    return handlerHelperCtx(cl, context.Background(), payload, url, method, authToken)
}
