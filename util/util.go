package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
)

func GenAuthorizationToken(consumerKey, consumerSecret string) string {
	// combine both consumer key & secret
	// base64 encode it to come up with a Authorization token
	key := []byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret))
	token := base64.URLEncoding.EncodeToString(key)

	return token
}

func BaseUrlBuilder(environment string) string {
	// build the base URL based on the code running environment
	// for non existing environments use the default base URL
	var baseURL string

	switch environment {
	case "sandbox":
		baseURL = "https://sandbox.safaricom.co.ke"
	case "production":
		baseURL = "https://sandbox.safaricom.co.ke"
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
