package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBusinessConsumer(t *testing.T) {
	testCases := []struct {
		name         string
		amount       float64
		txnType      txnType
		remarks      string
		customerNo   string
		qeueuTimeURL string
		resultURL    string
		check        func(t *testing.T, data *DarajaResParams, err error)
	}{
		{
			name:         "valid business consumer txn",
			amount:       10,
			txnType:      BusinessPayment,
			remarks:      "Business Payment Remarks",
			customerNo:   "0728762287",
			qeueuTimeURL: "https://mydomain.com/b2c/queue",
			resultURL:    "https://mydomain.com/b2c/result",
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

			data, err := client.BusinessToConsumer(tc.amount, tc.txnType, tc.customerNo, tc.remarks, tc.qeueuTimeURL, tc.resultURL)
			tc.check(t, data, err)
		})
	}
}
