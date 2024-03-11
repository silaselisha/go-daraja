package handler

type DarajaAuth interface {
	MpesaExpress(description, phoneNumber string, amount float64) (*StkCallback, error)
}

type Client struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RequestID    string `json:"requestId,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type StkCallback struct {
	MerchantRequestID string           `json:"merchant_request_id"`
	CheckoutRequestID string           `json:"checkout_request_id"`
	ResultCode        int64            `json:"result_code"`
	ResultDesc        string           `json:"result_desc"`
	CallbackMetadata  CallbackMetadata `json:"callback_metadata"`
}

type CallbackMetadata struct {
	Item []Item `json:"item"`
}

type Item struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExpressReqParams struct {
	BusinessShortCode string `json:"business_short_code"`
	Password          string `json:"password"`
	Timestamp         string `json:"timestamp"`
	TransactionType   string `json:"transaction_type"`
	Amount            string `json:"amount"`
	PartyA            string `json:"party_a"`
	PartyB            string `json:"party_b"`
	PhoneNumber       string `json:"phone_number"`
	CallBackURL       string `json:"call_back_url"`
	AccountReference  string `json:"account_reference"`
	TransactionDesc   string `json:"transaction_desc"`
}
