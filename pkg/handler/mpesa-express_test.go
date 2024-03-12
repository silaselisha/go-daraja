package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/silaselisha/go-daraja/util"
	"github.com/stretchr/testify/require"
)

func TestMpesaExpress(t *testing.T) {
	testCases := []struct {
		name        string
		phoneNumber string
		description string
		amount      float64
		check       func(t *testing.T, data []byte, err error)
	}{
		{
			name:        "valid transaction",
			phoneNumber: "0708374148",
			description: "test payment",
			amount:      1,
			check: func(t *testing.T, data []byte, err error) {
				
				var payload StkCallback
				require.NoError(t, err)
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)
				require.NotEmpty(t, payload)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Success. Request accepted for processing", payload.ResponseDescription)
			},
		},
		{
			name:        "invalid mpesa [PHONE NUMBER] transaction",
			phoneNumber: "070837414",
			description: "test payment",
			amount:      1,
			check: func(t *testing.T, data []byte, err error) {
				
				require.Error(t, err)
			},
		},
		{
			name:        "invalid [Daraja STK EXPRESS AMOUNT] transaction",
			phoneNumber: "0708374148",
			description: "test payment",
			amount:      4.99,
			check: func(t *testing.T, data []byte, err error) {

				require.NoError(t, err)
				var payload DarajaError
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)
				require.NotEmpty(t, payload)
				require.Equal(t, "400.002.02", payload.ErrorCode)
				require.Equal(t, "Bad Request - Invalid Amount", payload.ErrorMessage)
			},
		},
		{
			name:        "invalid [Daraja STK EXPRESS AMOUNT] transaction",
			phoneNumber: "0708374148",
			description: "test payment",
			amount:      1,
			check: func(t *testing.T, data []byte, err error) {
				fmt.Print(string(data))
				require.NoError(t, err)
				var payload DarajaError
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)
				require.NotEmpty(t, payload)
				require.Equal(t, "500.001.1001", payload.ErrorCode)
				require.Equal(t, "Unable to lock subscriber, a transaction is already in process for the current subscriber", payload.ErrorMessage)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			authToken := util.GenAuthorizationToken(testEnvs.DarajaConsumerKey, testEnvs.DarajaConsumerSecret)
			URL := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(testEnvs.DarajaEnvironment), "/oauth/v1/generate?grant_type=client_credentials")
			auth, err := NewDarajaAuth(URL, authToken)
			require.NoError(t, err)

			payload, err := auth.MpesaExpress(test.description, test.phoneNumber, test.amount)
			test.check(t, payload, err)
		})
	}
}
