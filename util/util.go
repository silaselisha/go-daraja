package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"time"
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

func GenTimestamp() string {
	// YYYYMMDDHHmmss
	timestamp := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
}

func GenSecurityCred(config *Configs, filePath string) (string, error) {
	passwordBuff := []byte(config.DarajaInitiatorPassword)
	fileName := "sandbox"
	if config.DarajaEnvironment == "production" {
		fileName = "production"
	}

	file := path.Join(filePath, "cert/", fileName+".cer")

	buff, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	// 1 define a cert block
	certBlock, _ := pem.Decode(buff)
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return "", err
	}
	// extract public key from the certificate
	key := cert.PublicKey.(*rsa.PublicKey)

	// encrypt password using RSA
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, key, passwordBuff)
	if err != nil {
		return "", err
	}

	securityCred := base64.StdEncoding.EncodeToString(cipherText)
	return string(securityCred), nil
}
