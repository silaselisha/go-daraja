package x509

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path"

	"github.com/silaselisha/go-daraja/pkg/internal/config"
)

func GenSecurityCred(config *config.Configs, filePath string) (string, error) {
	passwordBuff := []byte(config.DarajaInitiatorPassword)
	fileName := "sandbox"
	if config.MpesaEnvironment == "production" {
		fileName = "production"
	}

	file := path.Join(filePath, "cert", fileName+".cer")
	fmt.Println(file)
	buff, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// define a cert block
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
	fmt.Println(string(securityCred))
	return string(securityCred), nil
}
