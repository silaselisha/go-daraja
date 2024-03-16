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
	result := []byte(fmt.Sprintf("%s%s%s", util.Envs.DarajaBusinessShortCode, util.Envs.DarajaPassKey, util.Envs.DarajaTimestamp))
	password := base64.URLEncoding.EncodeToString(result)

	mobileNumber, err := util.PhoneNumberFormatter(phoneNumber)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	payload := ExpressReqParams{
		BusinessShortCode: util.Envs.DarajaBusinessShortCode,
		Password:          password,
		Timestamp:         util.Envs.DarajaTimestamp,
		TransactionType:   util.Envs.DarajaTransactionType,
		Amount:            amount,
		PartyA:            mobileNumber,
		PartyB:            util.Envs.DarajaPartyB,
		PhoneNumber:       mobileNumber,
		CallBackURL:       util.Envs.DarajaCallBackURL,
		AccountReference:  util.Envs.DarajaAccountRef,
		TransactionDesc:   description,
	}

	baseURL := util.BaseUrlBuilder(util.Envs.DarajaEnvironment)
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