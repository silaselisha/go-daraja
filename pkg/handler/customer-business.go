package handler

import (
	"fmt"
	"net/http"

	"github.com/silaselisha/go-daraja/internal/builder"
)

const (
	CANCELLED = iota
	COMPLETED
)

type C2BReqParams struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

func (cl *DarajaClient) CustomerToBusiness(authToken, confirmationURL, validationURL, responseType string) ([]byte, error) {
	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.DarajaEnvironment), "mpesa/c2b/v1/registerurl")

	payload := C2BReqParams{
		ShortCode:       cl.configs.DarajaBusinessShortCode,
		ResponseType:    responseType,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

	return handlerHelper[C2BReqParams](payload, URL, http.MethodPost, authToken)
}
