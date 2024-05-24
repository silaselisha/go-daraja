package handler

import (
	"fmt"
	"net/http"

	"github.com/silaselisha/go-daraja/pkg/internal/builder"
)

type b2cType int

const (
	CANCELLED b2cType = iota
	COMPLETED
)

type C2BReqParams struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

func (cl *DarajaClient) CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*DarajaResParams, error) {
	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/c2b/v1/registerurl")

	var command string
	switch {
	case responseType == 0:
		command = "CANCELLED"
	case responseType == 1:
		command = "COMPLETED"

	default:
		command = "CANCELLED"
	}
	payload := C2BReqParams{
		ShortCode:       cl.configs.DarajaBusinessShortCode,
		ResponseType:    command,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

	return handlerHelper[C2BReqParams](cl.logger, payload, URL, http.MethodPost, cl.accessToken)
}
