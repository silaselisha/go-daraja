package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBusinessExpress(t *testing.T) {
	testCases := []struct {
		name        string
		paymentRef  string
		callbackURL string
		partnerName string
		receiver    string
		amount      float64
		check       func(t *testing.T, data *DarajaResParams, err error)
	}{
		{
			name:        "valid business express checkout",
			paymentRef:  "TestAccount",
			callbackURL: "https://mydomain.com/b2b-express-checkout/",
			partnerName: "Test",
			receiver:    "174379",
			amount:      10,
			check: func(t *testing.T, data *DarajaResParams, err error) {
				fmt.Printf("Business Express: %+v\n", data)
				require.NoError(t, err)

				require.Equal(t, "0", data.ResponseBody.Code)
				require.Equal(t, "USSD Initiated Successfully", data.ResponseBody.Status)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)

			data, err := client.BusinessExpressCheckout(tc.paymentRef, tc.callbackURL, tc.partnerName, tc.receiver, tc.amount)
			tc.check(t, data, err)
		})
	}
}
