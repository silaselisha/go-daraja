package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBusinessBuyGoods(t *testing.T) {
	testCases := []struct {
		name            string
		username        string
		shortCode       string
		commandID       string
		remarks         string
		resultURL       string
		qeueuTimeOutURL string
		receiverID      string
		senderID        string
		amount          float64
		accountRefrence string
		check           func(t *testing.T, data *DarajaResParams, err error)
	}{
		{
			name:            "valid business buy goods txn",
			username:        "API_Usename",
			shortCode:       "000000",
			commandID:       "BusinessPayBill",
			remarks:         "ok",
			resultURL:       "https://mydomain.com/b2b/result/",
			qeueuTimeOutURL: "https://mydomain.com/b2b/queue/",
			receiverID:      "4",
			senderID:        "4",
			amount:          239,
			accountRefrence: "353353",
			check: func(t *testing.T, data *DarajaResParams, err error) {
				require.NoError(t, err)

				require.Equal(t, "0", data.ResponseCode)
				require.Equal(t, "Accept the service request successfully.", data.ResponseDescription)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)

			data, err := client.BusinessBuyGoods(tc.amount, tc.username, tc.shortCode, tc.commandID, tc.remarks, tc.resultURL, tc.qeueuTimeOutURL, tc.receiverID, tc.senderID, tc.accountRefrence)
			tc.check(t, data, err)
		})
	}
}
