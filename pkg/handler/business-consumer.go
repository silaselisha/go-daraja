package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/util"
)

const (
	SalaryPayment = iota
	BusinessPayment
	PromotionalPayment
)

func (cl *DarajaClientParams) BusinessToConsumer(amount, customerNo, txnType, remarks, qeueuTimeOutURL, resultURL, authToken string) ([]byte, error) {
	baseURL := util.BaseUrlBuilder(cl.configs.DarajaEnvironment)
	URL := fmt.Sprintf("%s/%s", baseURL, "mpesa/b2c/v3/paymentrequest")
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	securityCred, err := util.GenSecurityCred(cl.configs, "./../..")
	if err != nil {
		return nil, err
	}

	//TODO: check the validity of the phone number
	payload := B2CReqParams{
		OriginatorConversationID: ID.String(),
		InitiatorName:            cl.configs.DarajaInitiatorName,
		Amount:                   amount,
		CommandID:                txnType,
		PartyA:                   cl.configs.DarajaPartyA,
		PartyB:                   customerNo, //254728762287
		Remarks:                  remarks,
		QueueTimeOutURL:          qeueuTimeOutURL,
		ResultURL:                resultURL,
		SecurityCredential:       securityCred,
	}

	buff, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(buff))
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
	
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
