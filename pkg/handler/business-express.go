package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/util"
)

type BExpressCheckoutParams struct {
	primaryShortCode  string
	receiverShortCode string
	amount            float64
	paymentRef        string
	callbackUrl       string
	partnerName       string
	RequestRefID      string
}

func (cl *DarajaClientParams) BusinessExpressCheckout(authToken, paymentRef, callbackURL, partnerName, receiver string, amount float64) ([]byte, error) {
	URL := fmt.Sprintf("%s%s", util.BaseUrlBuilder(cl.configs.DarajaEnvironment), "/v1/ussdpush/get-msisdn")

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := BExpressCheckoutParams{
		primaryShortCode:  cl.configs.DarajaBusinessShortCode,
		receiverShortCode: receiver,
		amount:            amount,
		paymentRef:        paymentRef,
		callbackUrl:       callbackURL,
		partnerName:       partnerName,
		RequestRefID:      id.String(),
	}
	
	return handlerHelper(payload, URL, http.MethodPost, authToken)
}
