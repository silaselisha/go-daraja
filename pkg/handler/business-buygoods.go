package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	payload := &B2BReqParams{
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

	buff, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(buff))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
