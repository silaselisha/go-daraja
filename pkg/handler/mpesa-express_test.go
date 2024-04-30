package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMpesaExpress(t *testing.T) {
	testCases := []struct {
		name            string
		phoneNumber     string
		description     string
		transactionType string
		amount          float64
		consumerKey     string
		consumerSecret  string
		check           func(t *testing.T, data []byte, err error)
	}{
		{
			name:        "valid transaction",
			phoneNumber: "0708374149",
			description: "test payment",
			amount:      1,
			check: func(t *testing.T, data []byte, err error) {
				var payload NICallbackParams
				require.NoError(t, err)
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)

				require.NotEmpty(t, payload)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Success. Request accepted for processing", payload.ResponseDescription)
				require.Equal(t, "Success. Request accepted for processing", payload.CustomerMessage)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			payload, err := client.NIPush(tc.description, tc.phoneNumber, tc.amount)
			fmt.Println(string(payload))
			tc.check(t, payload, err)
		})
	}
}
