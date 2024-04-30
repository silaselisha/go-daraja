package handler

import (
	"fmt"
	"net/http"

	"github.com/silaselisha/go-daraja/internal/builder"
)

type TRX_CODE int

const (
	BG TRX_CODE = iota
	WA
	PB
	SM
	SB
)

type QRReqParams struct {
	MerchantName string
	RefNo        string
	Amount       float64
	TrxCode      string
	CPI          string
	Size         int64
}

type QRRespParams struct {
	ResponseCode        string
	RequestID           string
	ResponseDescription string
	QRCode              string
}

func (cl *DarajaClient) DynamicQRCode(amount float64, qrSize int64, trxCode TRX_CODE, refNo, merchantName string) ([]byte, error) {
	URL := fmt.Sprintf("%s/%s", builder.BaseUrlBuilder(cl.configs.DarajaEnvironment), "mpesa/qrcode/v1/generate")

	var txCode string
	switch {
	case trxCode == 0:
		txCode = "BG"
	case trxCode == 1:
		txCode = "WA"
	case trxCode == 2:
		txCode = "PB"
	case trxCode == 3:
		txCode = "SM"
	case trxCode == 4:
		txCode = "SB"
	}
	payload := QRReqParams{
		MerchantName: merchantName,
		RefNo:        refNo,
		Amount:       amount,
		TrxCode:      txCode,
		CPI:          cl.configs.DarajaCPI,
		Size:         qrSize,
	}

	return handlerHelper[QRReqParams](payload, URL, http.MethodPost, cl.accessToken)
}
