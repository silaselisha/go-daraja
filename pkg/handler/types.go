package handler

import (
	"fmt"

	"github.com/silaselisha/go-daraja/util"
)

type Daraja interface {
	ClientAuth() (*DarajaAuth, error)
	NIPush(description string, phoneNumber string, amount float64, authToken string) ([]byte, error)
	BusinessToConsumer(amount, customerNo, txnType, remarks, timeoutURL, resultURL, authToken string) ([]byte, error)
	CustomerToBusiness(authToken, confirmationURL, validationURL, responseType string) ([]byte, error)
	BusinessBuyGoods(amount float64, authToken, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountRefrence string) ([]byte, error)
}

type DarajaClientParams struct {
	configs *util.Configs
}

func NewDarajaClient(path string) (Daraja, error) {
	configs, err := util.LoadConfigs(path)
	fmt.Println(configs.DarajaEnvironment)
	if err != nil {
		return nil, err
	}
	return &DarajaClientParams{
		configs: configs,
	}, nil
}

type BusinessResParams struct {
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
