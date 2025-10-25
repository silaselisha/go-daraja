package handler

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

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
    client := &http.Client{}
    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cfgs.MpesaEnvironment), "oauth/v1/generate?grant_type=client_credentials")

    req, err := http.NewRequest(http.MethodGet, URL, nil)
    if err != nil {
        return nil, err
    }

    authToken := auth.GenAuthorizationToken(cfgs.DarajaConsumerKey, cfgs.DarajaConsumerSecret)

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Basic "+authToken)
    res, err := client.Do(req)
    if err != nil {
        // Network issue: return structured unreachable error without failing construction
        return &DarajaAuth{
            ErrorCode:    "500.003.1001",
            ErrorMessage: "Service is currently unreachable. Please try again later.",
        }, nil
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return &DarajaAuth{
            ErrorCode:    "500.003.1001",
            ErrorMessage: "Service is currently unreachable. Please try again later.",
        }, nil
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
        return &DarajaAuth{
            RequestID:    errShape.RequestID,
            ErrorCode:    errShape.ErrorCode,
            ErrorMessage: errShape.ErrorMessage,
        }, nil
    }

    // Fallback for non-JSON responses (e.g., HTML)
    return &DarajaAuth{
        ErrorCode:    "500.003.1001",
        ErrorMessage: "Service is currently unreachable. Please try again later.",
    }, nil
}
