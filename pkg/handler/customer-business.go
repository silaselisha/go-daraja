package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

const (
	CANCELLED = iota
	COMPLETED
)

func (cl *DarajaClientParams) CustomerToBusiness(authToken, confirmationURL, validationURL, responseType string) ([]byte, error) {
	URL := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(cl.configs.DarajaEnvironment), "mpesa/c2b/v1/registerurl")

	payload := &C2BReqParams{
		ShortCode:       cl.configs.DarajaBusinessShortCode,
		ResponseType:    responseType,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

	buff, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}