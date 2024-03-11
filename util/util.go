package util

import (
	"encoding/base64"
	"fmt"
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
