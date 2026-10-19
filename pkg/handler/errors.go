package handler

import (
    "errors"
    "fmt"
)

// Synthetic code and message used when no usable response was received from the M-Pesa API
// (network failure, unreadable body, or a non-JSON body such as an HTML error page).
const (
    unreachableCode    = "500.003.1001"
    unreachableMessage = "Service is currently unreachable. Please try again later."
)

// ErrUnreachable is wrapped by APIError when the M-Pesa API could not be reached or returned
// a body that could not be decoded. Use errors.Is(err, ErrUnreachable) to detect it.
var ErrUnreachable = errors.New("m-pesa service unreachable")

// APIError represents an error returned by the M-Pesa API or networking layer.
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
    base := fmt.Sprintf("pokea api error: code=%s status=%d message=%s", e.Code, e.Status, e.Message)
    if e.Op != "" {
        base = e.Op + ": " + base
    }
    if e.Err != nil {
        return base + ": " + e.Err.Error()
    }
    return base
}

func (e *APIError) Unwrap() error { return e.Err }

// unreachableError builds an APIError for failures where no usable M-Pesa API payload was received.
func unreachableError(op string, status int, body []byte, cause error) *APIError {
    err := ErrUnreachable
    if cause != nil {
        err = fmt.Errorf("%w: %w", ErrUnreachable, cause)
    }
    return &APIError{
        Status:  status,
        Code:    unreachableCode,
        Message: unreachableMessage,
        RawBody: body,
        Op:      op,
        Err:     err,
    }
}

// Failed reports whether the response carries an M-Pesa API error.
func (r *PokeaResParams) Failed() bool {
    return r != nil && r.ErrorCode != ""
}

// AsError returns the M-Pesa API error carried by the response as an *APIError,
// or nil if the response does not represent a failure.
func (r *PokeaResParams) AsError() error {
    if !r.Failed() {
        return nil
    }
    return &APIError{
        Code:      r.ErrorCode,
        Message:   r.ErrorMessage,
        RequestID: r.RequestID,
    }
}
