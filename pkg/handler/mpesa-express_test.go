package handler

import (
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
		check           func(t *testing.T, data *DarajaResParams, err error)
	}{
		{
			name:        "valid transaction",
			phoneNumber: "0708374149",
			description: "test payment",
			amount:      1,
            check: func(t *testing.T, data *DarajaResParams, err error) {
                fmt.Printf("Mpesa Express: %+v\n", data)
                require.NoError(t, err)
                if data.ErrorCode == "500.003.1001" || data.ErrorMessage == "Service is currently unreachable. Please try again later." {
                    require.Equal(t, "Service is currently unreachable. Please try again later.", data.ErrorMessage)
                    return
                }
                require.Equal(t, "0", data.ResponseCode)
                require.Equal(t, "Success. Request accepted for processing", data.ResponseDescription)
                require.Equal(t, "Success. Request accepted for processing", data.CustomerMessage)
            },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			payload, err := client.NIPush(tc.description, tc.phoneNumber, tc.amount)
			tc.check(t, payload, err)
		})
	}
}
