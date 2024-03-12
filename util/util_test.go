package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthGenUtil(t *testing.T) {
	consumerKey := "AinW3mPD40202e78JRxqacXGU0YF3laj"
	consumerSecret := "gMsGecybWqGDpCDL"

	token := GenAuthorizationToken(consumerKey, consumerSecret)
	require.NotEmpty(t, token)
	fmt.Print(token)
}

func TestPhoneNumberFormat(t *testing.T) {
	testCases := []struct {
		name        string
		phoneNumber string
		check func(t *testing.T, phoneNumber string, err error)
	}{
		{
			name: "07 valid phone number",
			phoneNumber: "0721024268",
			check: func(t *testing.T, phoneNumber string, err error) {
				require.NoError(t, err)
				require.Equal(t, "254721024268", phoneNumber)
			},
		},
		{
			name: "01 valid phone number",
			phoneNumber: "0121024268",
			check: func(t *testing.T, phoneNumber string, err error) {
				require.NoError(t, err)
				require.Equal(t, "254121024268", phoneNumber)
			},
		},
		{
			name: "invalid phone number",
			phoneNumber: "072102426",
			check: func(t *testing.T, phoneNumber string, err error) {
				require.Error(t, err)
				require.Empty(t, phoneNumber)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			phoneNumber, err := PhoneNumberFormatter(test.phoneNumber)
			test.check(t, phoneNumber, err)
		})
	}
}
