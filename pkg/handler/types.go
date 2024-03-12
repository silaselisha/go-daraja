package handler

type DarajaAuth interface {
	MpesaExpress(description, phoneNumber string, amount float64) (*[]byte, error)
}

type Client struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type StkCallback struct {
	MerchantRequestID string           `json:"merchant_request_id,omitempty"`
	CheckoutRequestID string           `json:"checkout_request_id,omitempty"`
	ResultCode        int64            `json:"result_code,omitempty"`
	ResultDesc        string           `json:"result_desc,omitempty"`
	CallbackMetadata  CallbackMetadata `json:"callback_metadata,omitempty"`
	RequestID         string           `json:"requestId,omitempty"`
	ErrorCode         string           `json:"errorCode,omitempty"`
	ErrorMessage      string           `json:"errorMessage,omitempty"`
}

type CallbackMetadata struct {
	Item []Item `json:"item,omitempty"`
}

type Item struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ExpressReqParams struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            float64
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}
