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

func (cl *DarajaClientParams) BusinessToConsumer(amount float64, commandID, remarks, qeueuTimeOutURL, resultURL, occassion string) ([]byte, error) {
	baseURL := util.BaseUrlBuilder(cl.configs.DarajaEnvironment)
	URL := fmt.Sprintf("%s/%s", baseURL, "mpesa/b2c/v3/paymentrequest")
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := B2CReqParams{
		OriginatorConversationID: ID.String(),
		InitiatorName:            cl.configs.DarajaInitiatorName,
		Amount:                   amount,
		CommandID:                commandID,
		PartyA:                   cl.configs.DarajaPartyA,
		PartyB:                   cl.configs.DarajaPartyB,
		Remarks:                  remarks,
		QueueTimeOutURL:          qeueuTimeOutURL,
		ResultURL:                resultURL,
		Occassion:                occassion,
	}

	buff, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buff, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buff, nil
}
