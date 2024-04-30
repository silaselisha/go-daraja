package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/pkg/x509"
	"github.com/silaselisha/go-daraja/internal/builder"
)

const (
	SalaryPayment = iota
	BusinessPayment
	PromotionalPayment
)

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

func (cl *DarajaClient) BusinessToConsumer(amount, customerNo, txnType, remarks, qeueuTimeOutURL, resultURL, authToken string) ([]byte, error) {
	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.DarajaEnvironment), "mpesa/b2c/v3/paymentrequest")
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	securityCred, err := x509.GenSecurityCred(cl.configs, "./../x509")
	if err != nil {
		return nil, err
	}

	mobileNumber, err := builder.PhoneNumberFormatter(customerNo)
	if err != nil {
		return nil, err
	}

	payload := B2CReqParams{
		OriginatorConversationID: ID.String(),
		InitiatorName:            cl.configs.DarajaInitiatorName,
		Amount:                   amount,
		CommandID:                txnType,
		PartyA:                   cl.configs.DarajaBusinessConsumerPartyA,
		PartyB:                   mobileNumber,
		Remarks:                  remarks,
		QueueTimeOutURL:          qeueuTimeOutURL,
		ResultURL:                resultURL,
		SecurityCredential:       securityCred,
	}

	return handlerHelper[B2CReqParams](payload, URL, http.MethodPost, authToken)
}
