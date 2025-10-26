# API Reference

Public interface `Daraja` methods:

```go
// Construct with NewClient or NewDarajaClient
NewClient(opts ...Option) (*DarajaClient, error)
NewDarajaClient(path string) (Daraja, error)

// Operations
NIPush(description string, phoneNumber string, amount float64) (*DarajaResParams, error)
BusinessToConsumer(amount float64, txnType txnType, customerNo, remarks, timeoutURL, resultURL string) (*DarajaResParams, error)
CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*DarajaResParams, error)
BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountRefrence string) (*DarajaResParams, error)
BusinessExpressCheckout(paymentRef, callbackURL, partnerName, receiver string, amount float64) (*DarajaResParams, error)
```

## Types

```go
type txnType int
const (
  SalaryPayment txnType = iota
  BusinessPayment
  PromotionalPayment
)

type b2cType int
const (
  CANCELLED b2cType = iota
  COMPLETED
)

// Response envelope
type DarajaResParams struct {
  ConversationID           string
  OriginatorConversationID string
  ResponseCode             string
  ResponseDescription      string
  CustomerMessage          string
  ResponseBody             struct {
    Code   string `json:"code"`
    Status string `json:"status"`
  }
  DarajaErrorParams
}

type DarajaErrorParams struct {
  RequestID    string `json:"requestId"`
  ErrorCode    string `json:"errorCode"`
  ErrorMessage string `json:"errorMessage"`
}
```
