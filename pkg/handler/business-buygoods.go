package handler

import (
	"fmt"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

type B2BReqParams struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	RecieverIdentifierType string
	Amount                 float64
	PartyA                 string
	PartyB                 string
	AccountReference       string
	Requester              string
	Remarks                string
	QueueTimeOutURL        string
	ResultURL              string
}

func (cl *DarajaClientParams) BusinessBuyGoods(amount float64, authToken, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference string) ([]byte, error) {
	URL := fmt.Sprintf("%s/%s", util.BaseUrlBuilder(cl.configs.DarajaEnvironment), "mpesa/b2b/v1/paymentrequest")

	securityCred, err := util.GenSecurityCred(cl.configs, "./../..")
	if err != nil {
		return nil, err
	}

	payload := B2BReqParams{
		Initiator:              username,
		SecurityCredential:     securityCred,
		CommandID:              commandID,
		SenderIdentifierType:   receiverID,
		RecieverIdentifierType: senderID,
		Amount:                 amount,
		PartyA:                 cl.configs.DarajaBusinessShortCode,
		PartyB:                 shortCode,
		AccountReference:       accountReference,
		Remarks:                remarks,
		QueueTimeOutURL:        queueTimeOutURL,
		ResultURL:              resultURL,
	}

	data, err := handlerHelper[B2BReqParams](payload, URL, http.MethodPost, authToken)
	return data, err
}
