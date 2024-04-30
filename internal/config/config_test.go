package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigs(t *testing.T) {
	testCases := []struct {
		name  string
		path  string
		check func(t *testing.T, envs *Configs, err error)
	}{
		{
			name: "invalid env path",
			path: "200",
			check: func(t *testing.T, envs *Configs, err error) {
				require.Error(t, err)
				require.Empty(t, envs)
			},
		},
		{
			name: "valid env path",
			path: "./../../example",
			check: func(t *testing.T, envs *Configs, err error) {
				fmt.Println(envs)
				require.NoError(t, err)
				require.NotEmpty(t, envs)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			envs, err := LoadConfigs(test.path)
			test.check(t, envs, err)
		})
	}
}
