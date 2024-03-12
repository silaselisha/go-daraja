package handler

import (
	"log"
	"os"
	"testing"

	"github.com/silaselisha/go-daraja/util"

	"github.com/stretchr/testify/require"
)

var testEnvs *util.Configs

func TestMain(m *testing.M) {
	envs, err := util.LoadConfigs("./../..")
	if err != nil {
		log.Fatal(err)
	}

	testEnvs = envs
	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	testCases := []struct {
		name      string
		authToken string
		check     func(t *testing.T, res any, err error)
	}{
		{
			name:      "200 OK Auth Request",
			authToken: "Basic cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(DarajaAuth).(*Client)
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
		},
		{
			name:      "Incorrect authorization type",
			authToken: "Baerer   cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ==",
			check: func(t *testing.T, response any, err error) {
				res := response.(DarajaAuth).(*Client)
				require.NoError(t, err)
				require.Equal(t, "Invalid Authentication passed", res.ErrorMessage)
				require.Equal(t, "400.008.01", res.ErrorCode)
				require.NotEmpty(t, res)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			auth, err := NewDarajaAuth(test.authToken)
			test.check(t, auth, err)
		})
	}
}
