package config

import (
	"fmt"
	"os"
	"path/filepath"
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
			path: "./../../../example",
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

func TestLoadConfigsLegacyEnv(t *testing.T) {
	t.Run("legacy .env file", func(t *testing.T) {
		dir := t.TempDir()
		env := "DARAJA_CONSUMER_KEY=old-key\nDARAJA_CONSUMER_SECRET=old-secret\nDARAJA_PASS_KEY=old-pass\n"
		require.NoError(t, os.WriteFile(filepath.Join(dir, ".env"), []byte(env), 0o600))

		cfg, err := LoadConfigs(dir)
		require.NoError(t, err)
		require.Equal(t, "old-key", cfg.PokeaConsumerKey)
		require.Equal(t, "old-secret", cfg.PokeaConsumerSecret)
		require.Equal(t, "old-pass", cfg.PokeaPassKey)
		require.Equal(t, "CustomerPayBillOnline", cfg.PokeaTransactionType)
	})

	t.Run("legacy OS env", func(t *testing.T) {
		t.Setenv("DARAJA_CONSUMER_KEY", "env-key")
		t.Setenv("DARAJA_CONSUMER_SECRET", "env-secret")

		cfg, err := LoadConfigs(t.TempDir())
		require.NoError(t, err)
		require.Equal(t, "env-key", cfg.PokeaConsumerKey)
		require.Equal(t, "env-secret", cfg.PokeaConsumerSecret)
	})

	t.Run("POKEA_* takes precedence", func(t *testing.T) {
		t.Setenv("DARAJA_CONSUMER_KEY", "old-key")
		t.Setenv("POKEA_CONSUMER_KEY", "new-key")
		t.Setenv("POKEA_CONSUMER_SECRET", "new-secret")

		cfg, err := LoadConfigs(t.TempDir())
		require.NoError(t, err)
		require.Equal(t, "new-key", cfg.PokeaConsumerKey)
	})
}
