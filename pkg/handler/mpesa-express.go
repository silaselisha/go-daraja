package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

func (cl *Client) MpesaExpress(description, phoneNumber string, amount float64) ([]byte, error) {
	result := []byte(fmt.Sprintf("%s%s%s", cl.configs.DarajaBusinessShortCode, cl.configs.DarajaPassKey, cl.configs.DarajaTimestamp))
	password := base64.URLEncoding.EncodeToString(result)

	mobileNumber, err := util.PhoneNumberFormatter(phoneNumber)
	if err != nil {
		log.Print(err)
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
		fmt.Print(err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(reqData))
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + cl.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return resData, nil
}