package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
)

func NewDarajaAuth(url string, authToken string) (DarajaAuth, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

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

	var result *Client
	json.Unmarshal(body, &result)
	return result, nil
}
