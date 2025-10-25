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
    ShortCode       string `json:"ShortCode"`
    ResponseType    string `json:"ResponseType"`
    ConfirmationURL string `json:"ConfirmationURL"`
    ValidationURL   string `json:"ValidationURL"`
}

func (cl *DarajaClient) CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*DarajaResParams, error) {
    URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.MpesaEnvironment), "mpesa/c2b/v1/registerurl")

	var command string
    switch responseType {
    case CANCELLED:
        command = "Cancelled"
    case COMPLETED:
        command = "Completed"
    default:
        command = "Cancelled"
    }
	payload := C2BReqParams{
		ShortCode:       cl.configs.DarajaBusinessShortCode,
		ResponseType:    command,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

	return handlerHelper(payload, URL, http.MethodPost, cl.AccessToken)
}
