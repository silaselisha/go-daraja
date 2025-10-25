package handler

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"

    "github.com/silaselisha/go-daraja/pkg/internal/auth"
    "github.com/silaselisha/go-daraja/pkg/internal/builder"
    "github.com/silaselisha/go-daraja/pkg/internal/config"
)

type DarajaAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func ClientAuth(cfgs *config.Configs) (*DarajaAuth, error) {
    // Backward-compatible wrapper with default timeout
    client := &http.Client{Timeout: 15 * time.Second}
    return ClientAuthWithClient(context.Background(), cfgs, client)
}

// ClientAuthWithClient performs token acquisition using the provided context and HTTP client.
func ClientAuthWithClient(ctx context.Context, cfgs *config.Configs, client *http.Client) (*DarajaAuth, error) {
    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cfgs.MpesaEnvironment), "oauth/v1/generate?grant_type=client_credentials")

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
    if err != nil {
        return nil, err
    }

    authToken := auth.GenAuthorizationToken(cfgs.DarajaConsumerKey, cfgs.DarajaConsumerSecret)

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Basic "+authToken)
    res, err := client.Do(req)
    if err != nil {
        return nil, &APIError{Code: "500.003.1001", Message: "Service is currently unreachable. Please try again later.", Op: "auth"}
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, &APIError{Code: "500.003.1001", Message: "Service is currently unreachable. Please try again later.", Op: "auth-read"}
    }

    var tokenRes DarajaAuth
    if err := json.Unmarshal(body, &tokenRes); err == nil {
        return &tokenRes, nil
    }

    // Try decode known error fields
    var errShape struct {
        RequestID    string `json:"requestId"`
        ErrorCode    string `json:"errorCode"`
        ErrorMessage string `json:"errorMessage"`
    }
    if err := json.Unmarshal(body, &errShape); err == nil && (errShape.ErrorCode != "" || errShape.ErrorMessage != "") {
        return nil, &APIError{Code: errShape.ErrorCode, Message: errShape.ErrorMessage, Op: "auth-decode"}
    }

    // Fallback for non-JSON responses (e.g., HTML)
    return nil, &APIError{Code: "500.003.1001", Message: "Service is currently unreachable. Please try again later.", Op: "auth-decode"}
}
