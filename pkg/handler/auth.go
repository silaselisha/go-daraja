package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

func NewDarajaAuth(consumerKey, consumerSecret string) (DarajaAuth, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(util.Envs.DarajaEnvironment), "oauth/v1/generate?grant_type=client_credentials")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	authToken := util.GenAuthorizationToken(consumerKey, consumerSecret)

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
			fmt.Print("EOF error")
			return nil, err
		}
		return nil, err
	}

	var darajaAuth *Client
	json.Unmarshal(body, &darajaAuth)
	return darajaAuth, nil
}
