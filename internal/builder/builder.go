package builder

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

func BaseUrlBuilder(environment string) string {
	// build the base URL based on the code running environment
	// for non existing environments use the default base URL
	var baseURL string

	switch environment {
	case "sandbox":
		baseURL = "https://sandbox.safaricom.co.ke"
	case "production":
		baseURL = "https://api.safaricom.co.ke"
	default:
		baseURL = "https://sandbox.safaricom.co.ke"
	}
	return baseURL
}

func PhoneNumberFormatter(phoneNumber string) (string, error) {
	// format the phoneNumber to 2547XXXXXXXX
	re := regexp.MustCompile(`^(07|01)\d{8}$`)
	if !re.MatchString(phoneNumber) {
		err := errors.New("invalid phone number")
		return "", err
	}

	return fmt.Sprintf("%s%s", "254", phoneNumber[1:]), nil
}

func GenTimestamp() string {
	// YYYYMMDDHHmmss
	timestamp := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
}
