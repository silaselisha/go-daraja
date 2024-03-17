package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMpesaExpress(t *testing.T) {
	testCases := []struct {
		name           string
		phoneNumber    string
		description    string
		amount         float64
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, data []byte, err error)
	}{
		{
			name:           "valid transaction",
			phoneNumber:    "0708374148",
			description:    "test payment",
			amount:         1,
			check: func(t *testing.T, data []byte, err error) {

				var payload NICallbackParams
				require.NoError(t, err)
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)
				require.NotEmpty(t, payload)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Success. Request accepted for processing", payload.ResponseDescription)
			},
		},
		{
			name:           "invalid mpesa [PHONE NUMBER] transaction",
			phoneNumber:    "070837414",
			description:    "test payment",
			amount:         1,
			check: func(t *testing.T, data []byte, err error) {

				require.Error(t, err)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../..")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			auth, err := client.ClientAuth()
			require.NoError(t, err)
			require.NotEmpty(t, auth)

			payload, err := client.NIPush(test.description, test.phoneNumber, test.amount, auth.AccessToken)
			test.check(t, payload, err)
		})
	}
}
