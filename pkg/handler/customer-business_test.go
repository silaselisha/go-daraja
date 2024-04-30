package handler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomerToBusiness(t *testing.T) {
	testCases := []struct {
		name            string
		responseType    string
		validationURL   string
		confirmationURL string
		check           func(t *testing.T, buff []byte, err error)
	}{
		{
			name:            "valid customer to business txn",
			responseType:    "Canceled",
			validationURL:   "https://mydomain.com/validation",
			confirmationURL: "https://mydomain.com/confirmation",
			check: func(t *testing.T, buff []byte, err error) {
				require.NotEmpty(t, buff)
				require.NoError(t, err)
				var payload BusinessResParams
				err = json.Unmarshal(buff, &payload)
				require.NoError(t, err)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Success", payload.ResponseDescription)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)

			buff, err := client.CustomerToBusiness(tc.confirmationURL, tc.validationURL, tc.responseType)
			tc.check(t, buff, err)
		})
	}
}
