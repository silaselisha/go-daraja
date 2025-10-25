package x509

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "embed"
    "encoding/base64"
    "encoding/pem"
    "fmt"
    "io/fs"
    "os"
    "path"

    "github.com/silaselisha/go-daraja/pkg/internal/config"
)

//go:embed cert/*.cer
var embeddedCerts embed.FS

// GenSecurityCred generates the security credential using either embedded certs or a provided filesystem path.
// If filePath is empty, embedded certificates are used.
func GenSecurityCred(config *config.Configs, filePath string) (string, error) {
	passwordBuff := []byte(config.DarajaInitiatorPassword)
	fileName := "sandbox"
	if config.MpesaEnvironment == "production" {
		fileName = "production"
	}

    var (
        buff []byte
        err  error
    )
    if filePath == "" {
        // Prefer embedded certificate assets
        buff, err = fs.ReadFile(embeddedCerts, path.Join("cert", fileName+".cer"))
    } else {
        file := path.Join(filePath, "cert", fileName+".cer")
        buff, err = os.ReadFile(file)
    }
    if err != nil {
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
    return string(securityCred), nil
}
