package handler

import (
	"encoding/json"
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
		check           func(t *testing.T, buff []byte, err error)
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
			check: func(t *testing.T, buff []byte, err error) {
				require.NotEmpty(t, buff)
				require.NoError(t, err)

				var payload BusinessResParams
				err = json.Unmarshal(buff, &payload)
				require.NoError(t, err)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Accept the service request successfully.", payload.ResponseDescription)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)

			buff, err := client.BusinessBuyGoods(tc.amount, tc.username, tc.shortCode, tc.commandID, tc.remarks, tc.resultURL, tc.qeueuTimeOutURL, tc.receiverID, tc.senderID, tc.accountRefrence)
			tc.check(t, buff, err)
		})
	}
}
