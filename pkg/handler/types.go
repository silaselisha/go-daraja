package handler

import (
	"github.com/silaselisha/go-daraja/pkg/internal/config"
)

type Daraja interface {
	NIPush(description string, phoneNumber string, amount float64) (*DarajaResParams, error)
	BusinessToConsumer(amount float64, txnType txnType, customerNo, remarks, timeoutURL, resultURL string) (*DarajaResParams, error)
	CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*DarajaResParams, error)
	BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountRefrence string) (*DarajaResParams, error)
	BusinessExpressCheckout(paymentRef, callbackURL, partnerName, receiver string, amount float64) (*DarajaResParams, error)
}

type DarajaClient struct {
	configs     *config.Configs
	AccessToken string
}

type DarajaResParams struct {
	ConversationID           string
	OriginatorConversationID string
	ResponseCode             string
	ResponseDescription      string
	CustomerMessage          string
	ResponseBody             struct {
		Code   string `json:"code"`
		Status string `json:"status"`
	}
	DarajaErrorParams
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

func NewDarajaClient(path string) (Daraja, error) {
	configs, err := config.LoadConfigs(path)
	if err != nil {
		return nil, err
	}

	auth, err := ClientAuth(configs)
	if err != nil {
		return nil, err
	}

	return &DarajaClient{
		configs:     configs,
		AccessToken: auth.AccessToken,
	}, nil
}
