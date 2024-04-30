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
		check       func(t *testing.T, buff []byte, err error)
	}{
		{
			name:        "valid business express checkout",
			paymentRef:  "TestAccount",
			callbackURL: "https://mydomain.com/b2b-express-checkout/",
			partnerName: "Test",
			receiver:    "174379",
			amount:      10,
			check: func(t *testing.T, buff []byte, err error) {
				require.NoError(t, err)
				fmt.Print(string(buff))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewDarajaClient("./../../example")
			require.NoError(t, err)
		
			buff, err := client.BusinessExpressCheckout(tc.paymentRef, tc.callbackURL, tc.partnerName, tc.receiver, tc.amount)
			tc.check(t, buff, err)
		})
	}
}
