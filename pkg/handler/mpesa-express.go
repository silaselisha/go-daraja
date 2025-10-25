package handler

import (
    "encoding/base64"
    "fmt"
    "net/http"

    "github.com/silaselisha/go-daraja/pkg/internal/builder"
)

type NICallbackParams struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

type ExpressReqParams struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            float64
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}

func (cl *DarajaClient) NIPush(description string, phoneNumber string, amount float64) (*DarajaResParams, error) {
	timestamp := builder.GenTimestamp()
	result := []byte(fmt.Sprintf("%s%s%s", cl.configs.DarajaBusinessShortCode, cl.configs.DarajaPassKey, timestamp))
    // Per Daraja docs, use standard Base64 encoding
    password := base64.StdEncoding.EncodeToString(result)

	mobileNumber, err := builder.PhoneNumberFormatter(phoneNumber)
	if err != nil {
		return nil, err
	}

	payload := ExpressReqParams{
		BusinessShortCode: cl.configs.DarajaBusinessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   cl.configs.DarajaTransactionType,
		Amount:            amount,
		PartyA:            mobileNumber,
		PartyB:            cl.configs.DarajaBusinessShortCode,
		PhoneNumber:       mobileNumber,
		CallBackURL:       cl.configs.DarajaCallBackURL,
		AccountReference:  cl.configs.DarajaAccountRef,
		TransactionDesc:   description,
	}

	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/stkpush/v1/processrequest")
	return handlerHelper(payload, URL, http.MethodPost, cl.AccessToken)
}
