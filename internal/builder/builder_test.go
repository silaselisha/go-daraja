package builder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhoneNumberFormat(t *testing.T) {
	testCases := []struct {
		name        string
		phoneNumber string
		check       func(t *testing.T, phoneNumber string, err error)
	}{
		{
			name:        "07 valid phone number",
			phoneNumber: "0721024268",
			check: func(t *testing.T, phoneNumber string, err error) {
				require.NoError(t, err)
				require.Equal(t, "254721024268", phoneNumber)
			},
		},
		{
			name:        "01 valid phone number",
			phoneNumber: "0121024268",
			check: func(t *testing.T, phoneNumber string, err error) {
				require.NoError(t, err)
				require.Equal(t, "254121024268", phoneNumber)
			},
		},
		{
			name:        "invalid phone number",
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

func TestGenTimestamp(t *testing.T) {
	timestamp := GenTimestamp()
	require.NotNil(t, timestamp)
	require.Equal(t, 14, len(timestamp))
}

func TestBaseURL(t *testing.T) {
	testCases := []struct {
		name        string
		environment string
		check       func(t *testing.T, URL string)
	}{
		{
			name:        "sandbox environment",
			environment: "sandbox",
			check: func(t *testing.T, URL string) {
				require.NotEmpty(t, URL)
				require.Equal(t, "https://sandbox.safaricom.co.ke", URL)
			},
		},
		{
			name:        "production environment",
			environment: "production",
			check: func(t *testing.T, URL string) {
				require.NotEmpty(t, URL)
				require.Equal(t, "https://sandbox.safaricom.co.ke", URL)
			},
		},
		{
			name:        "trivial default environment",
			environment: "",
			check: func(t *testing.T, URL string) {
				require.NotEmpty(t, URL)
				require.Equal(t, "https://sandbox.safaricom.co.ke", URL)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			URL := BaseUrlBuilder(test.environment)
			test.check(t, URL)
		})
	}
}
