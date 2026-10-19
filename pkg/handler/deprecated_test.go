package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Deprecated names must stay interchangeable with the Pokea names until they are removed.
func TestDeprecatedAliases(t *testing.T) {
	var cl *DarajaClient = &PokeaClient{}
	var _ Daraja = cl
	var _ Pokea = cl

	var res *PokeaResParams = &DarajaResParams{PokeaErrorParams: DarajaErrorParams{ErrorCode: "400.002.02"}}
	require.True(t, res.Failed())

	var auth *PokeaAuth = &DarajaAuth{AccessToken: "tok"}
	require.Equal(t, "tok", auth.AccessToken)
}
