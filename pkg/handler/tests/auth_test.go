package handler

import (
	"fmt"
	"testing"

	"github.com/silaselisha/go-daraja/pkg/handler"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	configs, err := handler.NewDarajaClient("./../../../example")
	require.NoError(t, err)
	require.NotEmpty(t, configs)
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name           string
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, res any, err error)
	}{
		{
			name: "200 OK Auth Request",
			check: func(t *testing.T, response any, err error) {
				fmt.Print(response)
				res := response.(*handler.DarajaAuth)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := handler.NewDarajaClient("./../../../example")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			auth, err := client.ClientAuth()
			tc.check(t, auth, err)
		})
	}
}
