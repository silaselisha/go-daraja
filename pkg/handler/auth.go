package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silaselisha/go-daraja/internal/auth"
	"github.com/silaselisha/go-daraja/internal/config"
	"github.com/silaselisha/go-daraja/internal/builder"
)

type DarajaAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func ClientAuth(cfgs *config.Configs) (*DarajaAuth, error) {
	client := &http.Client{}

	fmt.Println(cfgs.DarajaEnvironment)

	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cfgs.DarajaEnvironment), "oauth/v1/generate?grant_type=client_credentials")

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	authToken := auth.GenAuthorizationToken(cfgs.DarajaConsumerKey, cfgs.DarajaConsumerSecret)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+authToken)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, err
	}

	var darajaAuth *DarajaAuth
	if err := json.Unmarshal(body, &darajaAuth); err != nil {
		return nil, err
	}
	return darajaAuth, nil
}
