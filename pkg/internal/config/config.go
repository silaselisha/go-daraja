package config

import (
    "errors"
    "fmt"
    "strings"

    "github.com/spf13/viper"
)

type Configs struct {
    MpesaEnvironment                      string `mapstructure:"MPESA_ENVIRONMENT"`
    DarajaConsumerKey                      string `mapstructure:"DARAJA_CONSUMER_KEY"`
    DarajaConsumerSecret                   string `mapstructure:"DARAJA_CONSUMER_SECRET"`
    DarajaBusinessShortCode                string `mapstructure:"DARAJA_BUSINESS_SHORT_CODE"`
    DarajaPassKey                          string `mapstructure:"DARAJA_PASS_KEY"`
    DarajaTransactionType                  string `mapstructure:"DARAJA_TRANSACTION_TYPE"`
    DarajaCallBackURL                      string `mapstructure:"DARAJA_CALL_BACK_URL"`
    DarajaPartyA                           string `mapstructure:"DARAJA_PARTY_A"`
    DarajaPartyB                           string `mapstructure:"DARAJA_PARTY_B"`
    DarajaAccountRef                       string `mapstructure:"DARAJA_ACCOUNT_REF"`
    DarajaInitiatorName                    string `mapstructure:"DARAJA_INITIATOR_NAME"`
    DarajaInitiatorPassword                string `mapstructure:"DARAJA_INITIATOR_PASSWORD"`
    DarajaBusinessConsumerPartyA           string `mapstructure:"DARAJA_BUSINESS_CONSUMR_PARTY_A"`
    DarajaCPI                              string `mapstructure:"DARAJA_CREDIT_PARTY_IDENTIFIER"`
    DarajaBusinessExpressCheckoutShortCode string `mapstructure:"DARAJA_BUSINESS_EXPRESS_CHECKOUT_SHORT_CODE"`
}

// ValidateAuth performs minimal validation necessary to authenticate against Daraja.
// It intentionally focuses on the values required to obtain an access token so that
// client construction can succeed for most use cases. Operation-specific validation
// should be performed by the respective service methods when needed.
func (c *Configs) ValidateAuth() error {
    if c.DarajaConsumerKey == "" || c.DarajaConsumerSecret == "" {
        return errors.New("missing DARAJA_CONSUMER_KEY or DARAJA_CONSUMER_SECRET")
    }
    allowed := map[string]bool{"sandbox": true, "production": true}
    if c.MpesaEnvironment == "" {
        c.MpesaEnvironment = "sandbox"
    }
    if !allowed[c.MpesaEnvironment] {
        return fmt.Errorf("invalid MPESA_ENVIRONMENT: %s", c.MpesaEnvironment)
    }
    return nil
}

// LoadConfigs reads configuration from the environment and, if present, a .env file
// in the provided path. The .env file is optional; when it is not found, OS env vars
// and defaults are used. Defaults favor 12-factor friendliness.
func LoadConfigs(path string) (*Configs, error) {
    v := viper.New()
    v.SetConfigType("env")
    if strings.TrimSpace(path) != "" {
        v.AddConfigPath(path)
    }
    v.SetConfigName(".env")

    // Environment variable behavior
    v.AutomaticEnv()
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    // Sensible defaults
    v.SetDefault("MPESA_ENVIRONMENT", "sandbox")
    v.SetDefault("DARAJA_TRANSACTION_TYPE", "CustomerPayBillOnline")

    // Read .env when available; ignore if missing
    if err := v.ReadInConfig(); err != nil {
        // If the file is missing, proceed with env vars only
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, err
        }
    }

    var configs Configs
    if err := v.Unmarshal(&configs); err != nil {
        return nil, err
    }

    // Backstop defaults when tags were not populated
    if configs.MpesaEnvironment == "" {
        configs.MpesaEnvironment = "sandbox"
    }
    if configs.DarajaTransactionType == "" {
        configs.DarajaTransactionType = "CustomerPayBillOnline"
    }

    if err := configs.ValidateAuth(); err != nil {
        return nil, err
    }
    return &configs, nil
}
