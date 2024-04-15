package handler

import "github.com/silaselisha/go-daraja/util"

type Daraja interface {
	ClientAuth() (*DarajaAuth, error)
	NIPush(description string, phoneNumber string, amount float64, authToken,transactionType string) ([]byte, error)
	BusinessToConsumer(amount, customerNo, txnType, remarks, timeoutURL, resultURL, authToken string) ([]byte, error)
	CustomerToBusiness(authToken, confirmationURL, validationURL, responseType string) ([]byte, error)
}

type DarajaClientParams struct {
	configs *util.Configs
}

func NewDarajaClient(path string) (Daraja, error) {
	configs, err := util.LoadConfigs(path)
	if err != nil {
		return nil, err
	}
	return &DarajaClientParams{
		configs: configs,
	}, nil
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

type BusinessCustomerParams struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
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

type B2CReqParams struct {
	OriginatorConversationID string
	InitiatorName            string
	SecurityCredential       string
	CommandID                string
	Amount                   string
	PartyA                   string
	PartyB                   string
	Remarks                  string
	QueueTimeOutURL          string
	ResultURL                string
	Occassion                string
}

type C2BReqParams struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}
