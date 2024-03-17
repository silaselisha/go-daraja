package handler

import "github.com/silaselisha/go-daraja/util"

type Daraja interface {
	ClientAuth() (*DarajaAuth, error)
	NIPush(description string, phoneNumber string, amount float64, authToken string) ([]byte, error)
}

type DarajaClientParams struct {
	configs *util.Configs
}

type DarajaAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type NICallbackParams struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

type DarajaErrorParams struct {
	RequestID    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type CallbackMetadata struct {
	Item []ItemParams `json:"item,omitempty"`
}

type ItemParams struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
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
