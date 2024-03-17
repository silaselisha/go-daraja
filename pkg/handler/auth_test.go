package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	configs, err := NewDarajaClient(".")
	require.Error(t, err)
	require.Empty(t, configs)
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name           string
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, res any, err error)
	}{
		{
			name:           "200 OK Auth Request",
			check: func(t *testing.T, response any, err error) {
				fmt.Print(response)
				res := response.(*DarajaAuth)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../..")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			auth, err := client.ClientAuth()
			test.check(t, auth, err)
		})
	}
}
