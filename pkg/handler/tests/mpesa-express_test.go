package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/silaselisha/go-daraja/pkg/handler"
	"github.com/stretchr/testify/require"
)

func TestMpesaExpress(t *testing.T) {
	testCases := []struct {
		name           string
		phoneNumber    string
		description    string
		amount         float64
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, data []byte, err error)
	}{
		{
			name:        "valid transaction",
			phoneNumber: "0708374148",
			description: "test payment",
			amount:      1,
			check: func(t *testing.T, data []byte, err error) {
				var payload handler.NICallbackParams
				require.NoError(t, err)
				err = json.Unmarshal(data, &payload)
				require.NoError(t, err)

				require.NotEmpty(t, payload)
				require.Equal(t, "0", payload.ResponseCode)
				require.Equal(t, "Success. Request accepted for processing", payload.ResponseDescription)
				require.Equal(t, "Success. Request accepted for processing", payload.CustomerMessage)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			client, err := handler.NewDarajaClient("./../../..")
			require.NoError(t, err)
			require.NotEmpty(t, client)

			auth, err := client.ClientAuth()
			log.Printf("auth token:%+v\n", auth.AccessToken)
			require.NoError(t, err)
			require.NotEmpty(t, auth)

			payload, err := client.NIPush(test.description, test.phoneNumber, test.amount, auth.AccessToken)
			fmt.Print(string(payload))
			test.check(t, payload, err)
		})
	}
}
