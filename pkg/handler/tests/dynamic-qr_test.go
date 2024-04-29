package handler

import (
	"fmt"
	"testing"

	"github.com/silaselisha/go-daraja/pkg/handler"
	"github.com/stretchr/testify/require"
)

func TestDynamicQRCode(t *testing.T) {
	testCases := []struct {
		name         string
		amount       float64
		qrSize       int64
		refNo        string
		merchantName string
		trxCode      handler.TRX_CODE
		check        func(t *testing.T, buff []byte, err error)
	}{
		{
			name:         "valid QR Code gen",
			amount:       1,
			qrSize:       300,
			refNo:        "Test-Invoice",
			merchantName: "Test-Supermarket",
			trxCode:      handler.BG,
			check: func(t *testing.T, buff []byte, err error) {
				require.NoError(t, err)
				fmt.Print(string(buff))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := handler.NewDarajaClient("./../../../example")
			require.NoError(t, err)
			auth, err := client.ClientAuth()
			require.NoError(t, err)
			buff, err := client.DynamicQRCode(tc.amount, tc.qrSize, tc.trxCode, tc.refNo, tc.merchantName, auth.AccessToken)
			tc.check(t, buff, err)
		})
	}
}
