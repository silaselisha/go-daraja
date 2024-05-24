package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/pkg/internal/builder"
	"github.com/silaselisha/go-daraja/pkg/internal/x509"
)

type txnType int

const (
	SalaryPayment txnType = iota
	BusinessPayment
	PromotionalPayment
)

type B2CReqParams struct {
	OriginatorConversationID string
	InitiatorName            string
	SecurityCredential       string
	CommandID                string
	Amount                   float64
	PartyA                   string
	PartyB                   string
	Remarks                  string
	QueueTimeOutURL          string
	ResultURL                string
	Occassion                string
}

func (cl *DarajaClient) BusinessToConsumer(amount float64, txnType txnType, customerNo, remarks, qeueuTimeOutURL, resultURL string) (*DarajaResParams, error) {
	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/b2c/v3/paymentrequest")
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	securityCred, err := x509.GenSecurityCred(cl.configs, "./../internal/x509")
	if err != nil {
		return nil, err
	}

	mobileNumber, err := builder.PhoneNumberFormatter(customerNo)
	if err != nil {
		return nil, err
	}

	var commandID string
	switch {
	case txnType == 0:
		commandID = "SalaryPayment"
	case txnType == 1:
		commandID = "BusinessPayment"
	case txnType == 2:
		commandID = "PromotionalPayment"
	default:
		commandID = "SalaryPayment"
	}
	payload := B2CReqParams{
		OriginatorConversationID: ID.String(),
		InitiatorName:            cl.configs.DarajaInitiatorName,
		Amount:                   amount,
		CommandID:                commandID,
		PartyA:                   cl.configs.DarajaBusinessConsumerPartyA,
		PartyB:                   mobileNumber,
		Remarks:                  remarks,
		QueueTimeOutURL:          qeueuTimeOutURL,
		ResultURL:                resultURL,
		SecurityCredential:       securityCred,
	}

	return handlerHelper[B2CReqParams](cl.logger, payload, URL, http.MethodPost, cl.accessToken)
}
