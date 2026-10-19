package handler

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/silaselisha/pokea/pkg/internal/auth"
    "github.com/silaselisha/pokea/pkg/internal/builder"
    "github.com/silaselisha/pokea/pkg/internal/config"
)

type PokeaAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func ClientAuth(cfgs *config.Configs) (*PokeaAuth, error) {
    // Backward-compatible wrapper with default timeout
    client := &http.Client{Timeout: 15 * time.Second}
    return ClientAuthWithClient(context.Background(), cfgs, client)
}

// ClientAuthWithClient performs token acquisition using the provided context and HTTP client.
//
// On failure it returns a non-nil *APIError. For backward compatibility the
// PokeaAuth is also returned with its error fields populated.
func ClientAuthWithClient(ctx context.Context, cfgs *config.Configs, client *http.Client) (*PokeaAuth, error) {
    const op = "pokea auth"

    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cfgs.MpesaEnvironment), "oauth/v1/generate?grant_type=client_credentials")

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
    if err != nil {
        return nil, err
    }

    authToken := auth.GenAuthorizationToken(cfgs.PokeaConsumerKey, cfgs.PokeaConsumerSecret)

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Basic "+authToken)
    res, err := client.Do(req)
    if err != nil {
        return failedAuth(unreachableError(op, 0, nil, err))
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return failedAuth(unreachableError(op, res.StatusCode, nil, err))
    }

    var tokenRes PokeaAuth
    if err := json.Unmarshal(body, &tokenRes); err != nil {
        // Non-JSON body (e.g., HTML error page)
        return failedAuth(unreachableError(op, res.StatusCode, body, err))
    }

    if tokenRes.ErrorCode != "" || tokenRes.AccessToken == "" || res.StatusCode < 200 || res.StatusCode > 299 {
        apiErr := &APIError{
            Status:    res.StatusCode,
            Code:      tokenRes.ErrorCode,
            Message:   tokenRes.ErrorMessage,
            RequestID: tokenRes.RequestID,
            RawBody:   body,
            Op:        op,
        }
        if apiErr.Message == "" {
            if res.StatusCode >= 200 && res.StatusCode <= 299 {
                apiErr.Message = "no access token in response"
            } else {
                apiErr.Message = http.StatusText(res.StatusCode)
            }
        }
        return &tokenRes, apiErr
    }

    return &tokenRes, nil
}

// failedAuth returns a PokeaAuth mirroring apiErr alongside the error itself.
func failedAuth(apiErr *APIError) (*PokeaAuth, error) {
    return &PokeaAuth{RequestID: apiErr.RequestID, ErrorCode: apiErr.Code, ErrorMessage: apiErr.Message}, apiErr
}
