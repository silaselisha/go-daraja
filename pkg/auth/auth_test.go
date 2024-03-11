package auth

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/silaselisha/go-daraja/util"

	"github.com/stretchr/testify/require"
)

var envs *util.Configs
var err error

func TestMain(m *testing.M) {
	envs, err = util.LoadConfigs("./../..")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name      string
		url       string
		authToken string
		check     func(t *testing.T, res any, err error)
	}{
		{
			name:      "200 OK Auth Request",
			url:       fmt.Sprintf("%s/%s", util.BaseUrlBuilder(envs.DarajaEnvironment), "/oauth/v1/generate?grant_type=client_credentials"),
			authToken: "Basic cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(*DarajaAuth)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
		{
			name:      "Incorrect authorization type",
			url:       fmt.Sprintf("%s/%s", util.BaseUrlBuilder(envs.DarajaEnvironment), "/oauth/v1/generate?grant_type=client_credentials"),
			authToken: "Baerer   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(*DarajaAuth)
				require.NoError(t, err)
				require.Equal(t, "Invalid Authentication passed", res.ErrorMessage)
				require.Equal(t, "400.008.01", res.ErrorCode)
				require.NotEmpty(t, res)
			},
		},
		{
			name:      "Incorrect grant type",
			url:       fmt.Sprintf("%s/%s", util.BaseUrlBuilder(envs.DarajaEnvironment), "/oauth/v1/generate?grant_type=client_credentialss"),
			authToken: "Basic   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(*DarajaAuth)
				require.NoError(t, err)
				require.Equal(t, "Invalid grant type passed", res.ErrorMessage)
				require.Equal(t, "400.008.02", res.ErrorCode)
				require.NotEmpty(t, res)
			},
		},
		{
			name:      "Incorrect sandbox.safaricom.co.ke URL",
			url:       "sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials",
			authToken: "Basic   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(*DarajaAuth)
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
