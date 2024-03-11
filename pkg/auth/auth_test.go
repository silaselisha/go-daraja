package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	testCases := []struct {
		name      string
		url       string
		authToken string
		check     func(t *testing.T, res any, err error)
	}{
		{
			name: "200 OK Auth Request",
			url: "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
			authToken: "Basic cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, auth any, err error) {
				res := auth.(*DarajaAuth)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
		{
			name: "Incorrect authorization type",
			url: "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
			authToken: "Baerer   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, auth any, err error) {
				res := auth.(*DarajaAuth)
				require.NoError(t, err)
				require.Equal(t, "Invalid Authentication passed", res.ErrorMessage)
				require.Equal(t, "400.008.01", res.ErrorCode)
				require.NotEmpty(t, res)
			},
		},
		{
			name: "Incorrect grant type",
			url: "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentialss",
			authToken: "Basic   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, auth any, err error) {
				res := auth.(*DarajaAuth)
				require.NoError(t, err)
				require.Equal(t, "Invalid grant type passed", res.ErrorMessage)
				require.Equal(t, "400.008.02", res.ErrorCode)
				require.NotEmpty(t, res)
			},
		},
		{
			name: "Incorrect sandbox.safaricom.co.ke URL",
			url: "sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
			authToken: "Basic   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, auth any, err error) {
				res := auth.(*DarajaAuth)
				fmt.Print(err)
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			auth, err := NewDarajaAuth(test.url, test.authToken)
			test.check(t, auth, err)
		})
	}
}
