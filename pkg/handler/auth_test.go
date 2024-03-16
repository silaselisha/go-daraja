package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	testCases := []struct {
		name           string
		consumerKey    string
		consumerSecret string
		check          func(t *testing.T, res any, err error)
	}{
		{
			name:           "200 OK Auth Request",
			consumerKey:    "JDG40OnpvvRgXhgoPZ9GhGCTm1WZ42geJ66pH1tHIwwo4MrR",
			consumerSecret: "yQcMx6pBUMVjZ90ILmA3QGJzf0m0l2gwhY45l9S3EzcLkH8xOPdqIaE7DQiX5xyO",
			check: func(t *testing.T, response any, err error) {
				fmt.Print(response)
				res := response.(DarajaAuth).(*Client)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			auth, err := NewDarajaAuth(test.consumerKey, test.consumerSecret)
			test.check(t, auth, err)
		})
	}
}
