package handler

import (
	"encoding/json"
	"testing"

	"github.com/silaselisha/go-daraja/pkg/handler"
	"github.com/stretchr/testify/require"
)

func TestBusinessConsumer(t *testing.T) {
	testCases := []struct {
		name         string
		amount       string
		commandID    string
		remarks      string
		qeueuTimeURL string
		resultURL    string
		check        func(t *testing.T, buff []byte, err error)
	}{
		{
			name:         "valid business consumer tx",
			amount:       "1",
			commandID:    "BusinessPayment",
			remarks:      "Business Payment Remarks",
			qeueuTimeURL: "https://mydomain.com/b2c/queue",
			resultURL:    "https://mydomain.com/b2c/result",
			check: func(t *testing.T, buff []byte, err error) {
				require.NoError(t, err)
				require.NotNil(t, buff)

				var payload handler.BusinessCustomerParams
				err = json.Unmarshal(buff, &payload)
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := handler.NewDarajaClient("./../../..")
			require.NoError(t, err)

			auth, err := client.ClientAuth()
			require.NoError(t, err)
			require.NotEmpty(t, auth)
			buff, err := client.BusinessToConsumer(tc.amount, tc.commandID, tc.remarks, tc.qeueuTimeURL, tc.resultURL, auth.AccessToken)
			tc.check(t, buff, err)
		})
	}
}