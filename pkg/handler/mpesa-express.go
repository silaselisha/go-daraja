package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

func (cl *DarajaClientParams) NIPush(description string, phoneNumber string, amount float64, authToken string) ([]byte, error) {
	result := []byte(fmt.Sprintf("%s%s%s", cl.configs.DarajaBusinessShortCode, cl.configs.DarajaPassKey, cl.configs.DarajaTimestamp))
	password := base64.URLEncoding.EncodeToString(result)

	mobileNumber, err := util.PhoneNumberFormatter(phoneNumber)
	if err != nil {
		return nil, err
	}

	payload := ExpressReqParams{
		BusinessShortCode: cl.configs.DarajaBusinessShortCode,
		Password:          password,
		Timestamp:         cl.configs.DarajaTimestamp,
		TransactionType:   cl.configs.DarajaTransactionType,
		Amount:            amount,
		PartyA:            mobileNumber,
		PartyB:            cl.configs.DarajaPartyB,
		PhoneNumber:       mobileNumber,
		CallBackURL:       cl.configs.DarajaCallBackURL,
		AccountReference:  cl.configs.DarajaAccountRef,
		TransactionDesc:   description,
	}

	baseURL := util.BaseUrlBuilder(cl.configs.DarajaEnvironment)
	URL := fmt.Sprintf("%s/%s", baseURL, "mpesa/stkpush/v1/processrequest")

	reqData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}
