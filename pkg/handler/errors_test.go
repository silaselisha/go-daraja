package handler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/silaselisha/pokea/pkg/internal/config"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func stubClient(status int, body string, err error) *http.Client {
	return &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if err != nil {
			return nil, err
		}
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
			Request:    r,
		}, nil
	})}
}

// skipIfPokeaUnavailable skips live sandbox tests when the M-Pesa API rejects or cannot serve the request.
func skipIfPokeaUnavailable(t *testing.T, err error) {
	t.Helper()
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		t.Skipf("m-pesa sandbox returned an error: %v", err)
	}
}

func TestHandlerHelperErrors(t *testing.T) {
	networkErr := errors.New("dial tcp: connection refused")

	testCases := []struct {
		name   string
		status int
		body   string
		err    error
		check  func(t *testing.T, res *PokeaResParams, err error)
	}{
		{
			name:   "success",
			status: http.StatusOK,
			body:   `{"ResponseCode":"0","ResponseDescription":"Success. Request accepted for processing"}`,
			check: func(t *testing.T, res *PokeaResParams, err error) {
				require.NoError(t, err)
				require.False(t, res.Failed())
				require.NoError(t, res.AsError())
				require.Equal(t, "0", res.ResponseCode)
			},
		},
		{
			name: "network error",
			err:  networkErr,
			check: func(t *testing.T, res *PokeaResParams, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.ErrorIs(t, err, ErrUnreachable)
				require.ErrorIs(t, err, networkErr)
				require.Equal(t, 0, apiErr.Status)
				// Response still mirrors the error for backward compatibility
				require.True(t, res.Failed())
				require.Equal(t, unreachableCode, res.ErrorCode)
			},
		},
		{
			name:   "pokea error payload",
			status: http.StatusBadRequest,
			body:   `{"requestId":"abc-123","errorCode":"400.002.02","errorMessage":"Bad Request - Invalid Amount"}`,
			check: func(t *testing.T, res *PokeaResParams, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.NotErrorIs(t, err, ErrUnreachable)
				require.Equal(t, http.StatusBadRequest, apiErr.Status)
				require.Equal(t, "400.002.02", apiErr.Code)
				require.Equal(t, "abc-123", apiErr.RequestID)
				require.Equal(t, "400.002.02", res.ErrorCode)
				require.Equal(t, "Bad Request - Invalid Amount", res.ErrorMessage)
			},
		},
		{
			name:   "non-2xx without error payload",
			status: http.StatusInternalServerError,
			body:   `{}`,
			check: func(t *testing.T, res *PokeaResParams, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.Equal(t, http.StatusInternalServerError, apiErr.Status)
				require.NotNil(t, res)
			},
		},
		{
			name:   "non-JSON body",
			status: http.StatusBadGateway,
			body:   `<html>502 Bad Gateway</html>`,
			check: func(t *testing.T, res *PokeaResParams, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.ErrorIs(t, err, ErrUnreachable)
				require.Equal(t, http.StatusBadGateway, apiErr.Status)
				require.Equal(t, unreachableCode, res.ErrorCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cl := &PokeaClient{httpClient: stubClient(tc.status, tc.body, tc.err)}
			res, err := handlerHelperCtx(cl, context.Background(), ExpressReqParams{}, "https://example.invalid", http.MethodPost, "token")
			tc.check(t, res, err)
		})
	}
}

func TestClientAuthErrors(t *testing.T) {
	cfg := &config.Configs{MpesaEnvironment: "sandbox", PokeaConsumerKey: "key", PokeaConsumerSecret: "secret"}

	testCases := []struct {
		name   string
		status int
		body   string
		err    error
		check  func(t *testing.T, res *PokeaAuth, err error)
	}{
		{
			name:   "success",
			status: http.StatusOK,
			body:   `{"access_token":"tok","expires_in":"3599"}`,
			check: func(t *testing.T, res *PokeaAuth, err error) {
				require.NoError(t, err)
				require.Equal(t, "tok", res.AccessToken)
			},
		},
		{
			name: "network error",
			err:  errors.New("timeout"),
			check: func(t *testing.T, res *PokeaAuth, err error) {
				require.ErrorIs(t, err, ErrUnreachable)
				require.Equal(t, unreachableCode, res.ErrorCode)
			},
		},
		{
			name:   "invalid credentials",
			status: http.StatusBadRequest,
			body:   `{"requestId":"r1","errorCode":"400.008.01","errorMessage":"Invalid Authentication passed"}`,
			check: func(t *testing.T, res *PokeaAuth, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.Equal(t, "400.008.01", apiErr.Code)
				require.Equal(t, "r1", apiErr.RequestID)
				require.Empty(t, res.AccessToken)
			},
		},
		{
			name:   "empty body with 200",
			status: http.StatusOK,
			body:   `{}`,
			check: func(t *testing.T, res *PokeaAuth, err error) {
				var apiErr *APIError
				require.ErrorAs(t, err, &apiErr)
				require.Equal(t, "no access token in response", apiErr.Message)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ClientAuthWithClient(context.Background(), cfg, stubClient(tc.status, tc.body, tc.err))
			tc.check(t, res, err)
		})
	}
}

func TestNewClientAuthFailure(t *testing.T) {
	cfg := config.Configs{MpesaEnvironment: "sandbox"}
	cl, err := NewClient(WithConfig(cfg), WithHTTPClient(stubClient(0, "", errors.New("offline"))))
	require.ErrorIs(t, err, ErrUnreachable)
	require.Nil(t, cl)
}
