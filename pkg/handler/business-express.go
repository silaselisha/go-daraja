package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/silaselisha/go-daraja/pkg/internal/builder"
)

type BExpressCheckoutParams struct {
	PrimaryShortCode  string  `json:"primaryShortCode"`
	ReceiverShortCode string  `json:"receiverShortCode"`
	Amount            float64 `json:"amount"`
	PaymentRef        string  `json:"paymentRef"`
	CallbackURL       string  `json:"callbackUrl"`
	PartnerName       string  `json:"partnerName"`
	RequestRefID      string  `json:"requestRefId"`
}

func (cl *DarajaClient) BusinessExpressCheckout(paymentRef, callbackURL, partnerName, receiver string, amount float64) (*DarajaResParams, error) {
	URL := fmt.Sprintf("%s%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "/v1/ussdpush/get-msisdn")

	requestRedID, err := uuid.NewRandom()
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
		RequestRefID:      requestRedID.String(),
	}

	fmt.Printf("PAYLOAD: %+v\n", payload)
	return handlerHelper(payload, URL, http.MethodPost, cl.AccessToken)
}
