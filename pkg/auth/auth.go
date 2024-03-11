package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DarajaAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

const (
	URL                 = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	AUTHORIZATION_TOKEN = "cFJZcjZ6anEwaThMMXp6d1FETUxwWkIzeVBDa2hNc2M6UmYyMkJmWm9nMHFRR2xWOQ=="
)

func NewDarajaAuth(url string, authToken string) (*DarajaAuth, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+authToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
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

	var result *DarajaAuth
	json.Unmarshal(body, &result)
	return result, nil
}
