package handler

import (
	"github.com/silaselisha/go-daraja/internal/config"
)

type Daraja interface {
	NIPush(description string, phoneNumber string, amount float64) ([]byte, error)
	BusinessToConsumer(amount float64, customerNo, txnType, remarks, timeoutURL, resultURL string) ([]byte, error)
	CustomerToBusiness(confirmationURL, validationURL, responseType string) ([]byte, error)
	BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountRefrence string) ([]byte, error)
	BusinessExpressCheckout(paymentRef, callbackURL, partnerName, receiver string, amount float64) ([]byte, error)
	DynamicQRCode(amount float64, qrSize int64, trxCode TRX_CODE, refNo, marchantName string) ([]byte, error)
}

type DarajaClient struct {
	configs     *config.Configs
	accessToken string
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
		accessToken: auth.AccessToken,
	}, nil
}
