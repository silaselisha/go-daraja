package handler

import (
	"testing"

	"github.com/silaselisha/pokea/pkg/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	configs, err := NewPokeaClient("./../../example")
	skipIfPokeaUnavailable(t, err)
	require.NoError(t, err)
	require.NotEmpty(t, configs)
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name           string
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, res any, err error)
	}{
		{
			name: "200 OK Auth Request",
			check: func(t *testing.T, response any, err error) {
				res := response.(*PokeaAuth)
				require.NotEmpty(t, res)
				skipIfPokeaUnavailable(t, err)
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			auth, err := ClientAuth(&config.Configs{MpesaEnvironment: "sandbox", PokeaConsumerKey: "JDG40OnpvvRgXhgoPZ9GhGCTm1WZ42geJ66pH1tHIwwo4MrR", PokeaConsumerSecret: "yQcMx6pBUMVjZ90ILmA3QGJzf0m0l2gwhY45l9S3EzcLkH8xOPdqIaE7DQiX5xyO"})
			tc.check(t, auth, err)
		})
	}
}
