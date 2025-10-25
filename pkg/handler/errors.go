package handler

import (
    "errors"
    "fmt"
)

// APIError represents an error returned by the Daraja API or networking layer.
// Status may be 0 for client-side/network errors where no HTTP response was received.
type APIError struct {
    Status    int
    Code      string
    Message   string
    RequestID string
    RawBody   []byte
    Op        string
    Err       error
}

func (e *APIError) Error() string {
    base := fmt.Sprintf("daraja api error: code=%s status=%d message=%s", e.Code, e.Status, e.Message)
    if e.Op != "" {
        base = e.Op + ": " + base
    }
    if e.Err != nil {
        return base + ": " + e.Err.Error()
    }
    return base
}

func (e *APIError) Unwrap() error { return e.Err }

// asDarajaError converts an error into a Daraja-like response struct for
// backward compatibility with the previous API that encoded errors in responses.
func asDarajaError(err error) *DarajaResParams {
    var apiErr *APIError
    if errors.As(err, &apiErr) {
        return &DarajaResParams{
            DarajaErrorParams: DarajaErrorParams{
                RequestID:    apiErr.RequestID,
                ErrorCode:    apiErr.Code,
                ErrorMessage: apiErr.Message,
            },
        }
    }
    return &DarajaResParams{
        DarajaErrorParams: DarajaErrorParams{
            ErrorCode:    "500.003.1001",
            ErrorMessage: "Service is currently unreachable. Please try again later.",
        },
    }
}
