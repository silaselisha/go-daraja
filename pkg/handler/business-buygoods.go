package handler

import (
    "context"
    "fmt"
    "net/http"

    "github.com/silaselisha/go-daraja/pkg/internal/x509"
    "github.com/silaselisha/go-daraja/pkg/internal/builder"
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

func (cl *DarajaClient) BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference string) (*DarajaResParams, error) {
    return cl.BusinessBuyGoodsCtx(context.Background(), amount, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference)
}

func (cl *DarajaClient) BusinessBuyGoodsCtx(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference string) (*DarajaResParams, error) {
    ctx := context.Background()
    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/b2b/v1/paymentrequest")

    // Use embedded certs by passing empty path; allow override later via options if needed
    securityCred, err := x509.GenSecurityCred(cl.configs, "")
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

    data, err := cl.handlerHelperCtx(ctx, payload, URL, http.MethodPost, cl.AccessToken)
	return data, err
}
