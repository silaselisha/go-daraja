package x509

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/silaselisha/go-daraja/pkg/internal/config"
)

func TestGenSecurityCred(t *testing.T) {
	cred, err := GenSecurityCred(&config.Configs{MpesaEnvironment: "sandbox"}, ".")
	require.NoError(t, err)
	require.NotNil(t, cred)
	fmt.Println(cred)
}
