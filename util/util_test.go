package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthGenUtil(t *testing.T) {
	consumerKey := "AinW3mPD40202e78JRxqacXGU0YF3laj"
	consumerSecret := "gMsGecybWqGDpCDL"

	token := GenAuthorizationToken(consumerKey, consumerSecret)
	require.NotEmpty(t, token)
	fmt.Print(token)
}
