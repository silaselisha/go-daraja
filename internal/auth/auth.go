package auth

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
