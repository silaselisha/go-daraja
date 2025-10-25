package handler

import (
    "net/http"
    "time"

    "github.com/silaselisha/go-daraja/pkg/internal/config"
)

type Daraja interface {
	NIPush(description string, phoneNumber string, amount float64) (*DarajaResParams, error)
	BusinessToConsumer(amount float64, txnType txnType, customerNo, remarks, timeoutURL, resultURL string) (*DarajaResParams, error)
	CustomerToBusiness(confirmationURL, validationURL string, responseType b2cType) (*DarajaResParams, error)
	BusinessBuyGoods(amount float64, username, shortCode, commandID, remarks, resultURL, queueTimeOutURL, receiverID, senderID, accountRefrence string) (*DarajaResParams, error)
	BusinessExpressCheckout(paymentRef, callbackURL, partnerName, receiver string, amount float64) (*DarajaResParams, error)
}

type DarajaClient struct {
	configs     *config.Configs
	AccessToken string
    httpClient  *http.Client
    logger      Logger
}

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

type CallbackMetadata struct {
	Item []ItemParams `json:"item,omitempty"`
}

type ItemParams struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Functional options for client construction
type Option func(*clientOptions)

type clientOptions struct {
    envFile    string
    cfg        *config.Configs
    httpClient *http.Client
    logger     Logger
}

// WithEnvFile specifies a directory containing a .env file to load.
// OS environment variables are always read; .env is optional.
func WithEnvFile(path string) Option {
    return func(o *clientOptions) { o.envFile = path }
}

// WithConfig provides an explicit configuration, bypassing env loading.
func WithConfig(cfg config.Configs) Option {
    return func(o *clientOptions) { o.cfg = &cfg }
}

// WithHTTPClient injects a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
    return func(o *clientOptions) { o.httpClient = hc }
}

// Logger is a minimal logging interface used by the client.
type Logger interface {
    Debugf(format string, args ...any)
    Infof(format string, args ...any)
    Errorf(format string, args ...any)
}

// WithLogger injects a logger implementation for observability.
func WithLogger(l Logger) Option {
    return func(o *clientOptions) { o.logger = l }
}

// NewClient constructs a Daraja client using functional options.
// Defaults: sandbox environment, 15s HTTP timeout, OS envs.
func NewClient(opts ...Option) (*DarajaClient, error) {
    // Gather options
    var o clientOptions
    for _, opt := range opts {
        opt(&o)
    }

    // Resolve configuration
    var cfg *config.Configs
    var err error
    if o.cfg != nil {
        cfg = o.cfg
    } else {
        cfg, err = config.LoadConfigs(o.envFile)
        if err != nil {
            return nil, err
        }
    }

    // Resolve HTTP client
    httpClient := o.httpClient
    if httpClient == nil {
        httpClient = &http.Client{Timeout: 15 * time.Second}
    }

    // Authenticate to fetch access token
    auth, err := ClientAuth(cfg)
    if err != nil {
        return nil, err
    }

    return &DarajaClient{
        configs:     cfg,
        AccessToken: auth.AccessToken,
        httpClient:  httpClient,
        logger:      o.logger,
    }, nil
}

// NewDarajaClient maintains backward compatibility for the existing API.
// It delegates to NewClient with an optional .env path.
func NewDarajaClient(path string) (Daraja, error) {
    cl, err := NewClient(WithEnvFile(path))
    if err != nil {
        return nil, err
    }
    return cl, nil
}
