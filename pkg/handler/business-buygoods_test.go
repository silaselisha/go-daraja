package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/silaselisha/pokea/pkg/internal/config"
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
		check           func(t *testing.T, data *PokeaResParams, err error)
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
			check: func(t *testing.T, data *PokeaResParams, err error) {
				skipIfPokeaUnavailable(t, err)
				require.NoError(t, err)

				require.Equal(t, "0", data.ResponseCode)
				require.Equal(t, "Accept the service request successfully.", data.ResponseDescription)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewPokeaClient("./../../example")
			skipIfPokeaUnavailable(t, err)
			require.NoError(t, err)

			data, err := client.BusinessBuyGoods(tc.amount, tc.username, tc.shortCode, tc.commandID, tc.remarks, tc.resultURL, tc.qeueuTimeOutURL, tc.receiverID, tc.senderID, tc.accountRefrence)
			tc.check(t, data, err)
		})
	}
}

func TestBusinessBuyGoodsIdentifierTypes(t *testing.T) {
	var body B2BReqParams
	hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"ResponseCode":"0"}`)), Header: make(http.Header), Request: r}, nil
	})}
	cl := &PokeaClient{configs: &config.Configs{MpesaEnvironment: "sandbox", PokeaInitiatorPassword: "secret"}, httpClient: hc}

	_, err := cl.BusinessBuyGoods(1, "apiuser", "000000", "BusinessBuyGoods", "remarks", "https://example.com/r", "https://example.com/t", "2", "4", "ref")
	require.NoError(t, err)
	require.Equal(t, "2", body.RecieverIdentifierType)
	require.Equal(t, "4", body.SenderIdentifierType)
}
