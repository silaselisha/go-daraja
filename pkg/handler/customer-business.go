package handler

import (
    "context"
    "fmt"
    "net/http"

    "github.com/silaselisha/pokea/pkg/internal/builder"
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

func (cl *PokeaClient) CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*PokeaResParams, error) {
    return cl.CustomerToBusinessCtx(context.Background(), confirmationURL, validationURL, responseType)
}

func (cl *PokeaClient) CustomerToBusinessCtx(ctx context.Context, confirmationURL, validationURL string, responseType b2cType) (*PokeaResParams, error) {
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
		ShortCode:       cl.configs.PokeaBusinessShortCode,
		ResponseType:    command,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

    return handlerHelperCtx(cl, ctx, payload, URL, http.MethodPost, cl.AccessToken)
}
