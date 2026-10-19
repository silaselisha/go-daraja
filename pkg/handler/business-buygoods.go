package handler

import (
    "context"
    "fmt"
    "net/http"

    "github.com/silaselisha/pokea/pkg/internal/x509"
    "github.com/silaselisha/pokea/pkg/internal/builder"
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

func (cl *PokeaClient) BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference string) (*PokeaResParams, error) {
    return cl.BusinessBuyGoodsCtx(context.Background(), amount, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference)
}

func (cl *PokeaClient) BusinessBuyGoodsCtx(ctx context.Context, amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountReference string) (*PokeaResParams, error) {
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
		SenderIdentifierType:   senderID,
		RecieverIdentifierType: receiverID,
		Amount:                 amount,
		PartyA:                 cl.configs.PokeaBusinessShortCode,
		PartyB:                 shortCode,
		AccountReference:       accountReference,
		Remarks:                remarks,
		QueueTimeOutURL:        queueTimeOutURL,
		ResultURL:              resultURL,
	}

    data, err := handlerHelperCtx(cl, ctx, payload, URL, http.MethodPost, cl.AccessToken)
	return data, err
}
