package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomerToBusiness(t *testing.T) {
	testCases := []struct {
		name            string
		responseType    b2cType
		validationURL   string
		confirmationURL string
		check           func(t *testing.T, data *DarajaResParams, err error)
	}{
		{
			name:            "valid customer to business cancelled txn",
			responseType:    COMPLETED,
			validationURL:   "https://mydomain.com/validation",
			confirmationURL: "https://mydomain.com/confirmation",
			check: func(t *testing.T, data *DarajaResParams, err error) {
				require.NoError(t, err)

				require.Equal(t, "0", data.ResponseCode)
				require.Equal(t, "Success", data.ResponseDescription)
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
