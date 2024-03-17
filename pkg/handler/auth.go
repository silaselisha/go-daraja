package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

func NewDarajaClient(path string) (Daraja, error) {
	configs, err := util.LoadConfigs(path)
	if err != nil {
		return nil, err
	}
	return &DarajaClientParams{
		configs: configs,
	}, nil
}

func (cl *DarajaClientParams) ClientAuth() (*DarajaAuth, error) {
	client := &http.Client{}

	configs, err := util.LoadConfigs("./../..")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(configs.DarajaEnvironment), "oauth/v1/generate?grant_type=client_credentials")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	authToken := util.GenAuthorizationToken(cl.configs.DarajaConsumerKey, cl.configs.DarajaConsumerSecret)

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
	json.Unmarshal(body, &darajaAuth)
	return darajaAuth, nil
}
