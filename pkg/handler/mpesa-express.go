package handler

import (
	"fmt"
	"net/http"

	"github.com/silaselisha/go-daraja/util"
)

func (cl *Client) MpesaExpress(description, phoneNumber string, amount float64) (*StkCallback, error){
	_, err := util.PhoneNumberFormatter(phoneNumber)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	envs, err := util.LoadConfigs("./../..")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	baseURL := util.BaseUrlBuilder(envs.DarajaEnvironment)
	URL := fmt.Sprintf("%s/%s", baseURL, "mpesa/stkpush/v1/processrequest")

	_, err = http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return nil, nil
}