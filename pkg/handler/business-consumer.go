package handler

import (
    "fmt"
    "net/http"

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
    OriginatorConversationID string  `json:"OriginatorConversationID,omitempty"`
    InitiatorName      string  `json:"InitiatorName"`
    SecurityCredential string  `json:"SecurityCredential"`
    CommandID          string  `json:"CommandID"`
    Amount             float64 `json:"Amount"`
    PartyA             string  `json:"PartyA"`
    PartyB             string  `json:"PartyB"`
    Remarks            string  `json:"Remarks"`
    QueueTimeOutURL    string  `json:"QueueTimeOutURL"`
    ResultURL          string  `json:"ResultURL"`
    Occasion           string  `json:"Occasion,omitempty"`
}

func (cl *DarajaClient) BusinessToConsumer(amount float64, txnType txnType, customerNo, remarks, qeueuTimeOutURL, resultURL string) (*DarajaResParams, error) {
    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/b2c/v1/paymentrequest")

    // OriginatorConversationID is optional for v1; include when available to align with docs
    // We can reuse the SecurityCredential entropy to derive a pseudo-unique ID if none
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
        InitiatorName:      cl.configs.DarajaInitiatorName,
        Amount:             amount,
        CommandID:          commandID,
        PartyA:             cl.configs.DarajaBusinessConsumerPartyA,
        PartyB:             mobileNumber,
        Remarks:            remarks,
        QueueTimeOutURL:    qeueuTimeOutURL,
        ResultURL:          resultURL,
        SecurityCredential: securityCred,
    }

    return handlerHelper(payload, URL, http.MethodPost, cl.AccessToken)
}
