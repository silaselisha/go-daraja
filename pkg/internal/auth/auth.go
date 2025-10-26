package auth

import (
    "encoding/base64"
    "fmt"
)

func GenAuthorizationToken(consumerKey, consumerSecret string) string {
    // combine both consumer key & secret and Base64-encode per HTTP Basic Auth spec
    key := []byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret))
    token := base64.StdEncoding.EncodeToString(key)
    return token
}
