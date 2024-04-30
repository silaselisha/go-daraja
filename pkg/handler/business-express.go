package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/internal/builder"
)

type BExpressCheckoutParams struct {
	PrimaryShortCode  string  `json:"primaryShortCode"`
	ReceiverShortCode string   `json:"receiverShortCode"`
	Amount            float64 `json:"amount"`
	PaymentRef        string  `json:"paymentRef"`
	CallbackURL       string  `json:"callbackUrl"`
	PartnerName       string  `json:"partnerName"`
	RequestRefID      string  `json:"RequestRefID"`
}

func (cl *DarajaClient) BusinessExpressCheckout(authToken, paymentRef, callbackURL, partnerName, receiver string, amount float64) ([]byte, error) {
	URL := fmt.Sprintf("%s%s", builder.BaseUrlBuilder(cl.configs.DarajaEnvironment), "/v1/ussdpush/get-msisdn")

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := BExpressCheckoutParams{
		PrimaryShortCode:  cl.configs.DarajaBusinessExpressCheckoutShortCode,
		ReceiverShortCode: receiver,
		Amount:            amount,
		PaymentRef:        paymentRef,
		CallbackURL:       callbackURL,
		PartnerName:       partnerName,
		RequestRefID:      id.String(),
	}

	fmt.Println(payload)
	return handlerHelper(payload, URL, http.MethodPost, authToken)
}
