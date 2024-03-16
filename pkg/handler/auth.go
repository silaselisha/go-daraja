package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/silaselisha/go-daraja/util"
)

func NewDarajaAuth(consumerKey, consumerSecret string) (DarajaAuth, error) {
	client := &http.Client{}
	envs, err := util.LoadConfigs(os.Getenv(".env"))
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(envs.DarajaEnvironment), "oauth/v1/generate?grant_type=client_credentials")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	authToken := util.GenAuthorizationToken(consumerKey, consumerSecret)
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic " + authToken)
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

	var darajaAuth *Client = &Client{
		config: envs,
	}
	json.Unmarshal(body, &darajaAuth)
	return darajaAuth, nil
}
