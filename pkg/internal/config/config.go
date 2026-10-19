package config

import (
    "errors"
    "fmt"
    "os"
    "reflect"
    "strings"

    "github.com/spf13/viper"
)

type Configs struct {
    MpesaEnvironment                      string `mapstructure:"MPESA_ENVIRONMENT"`
    PokeaConsumerKey                      string `mapstructure:"POKEA_CONSUMER_KEY"`
    PokeaConsumerSecret                   string `mapstructure:"POKEA_CONSUMER_SECRET"`
    PokeaBusinessShortCode                string `mapstructure:"POKEA_BUSINESS_SHORT_CODE"`
    PokeaPassKey                          string `mapstructure:"POKEA_PASS_KEY"`
    PokeaTransactionType                  string `mapstructure:"POKEA_TRANSACTION_TYPE"`
    PokeaCallBackURL                      string `mapstructure:"POKEA_CALL_BACK_URL"`
    PokeaPartyA                           string `mapstructure:"POKEA_PARTY_A"`
    PokeaPartyB                           string `mapstructure:"POKEA_PARTY_B"`
    PokeaAccountRef                       string `mapstructure:"POKEA_ACCOUNT_REF"`
    PokeaInitiatorName                    string `mapstructure:"POKEA_INITIATOR_NAME"`
    PokeaInitiatorPassword                string `mapstructure:"POKEA_INITIATOR_PASSWORD"`
    PokeaBusinessConsumerPartyA           string `mapstructure:"POKEA_BUSINESS_CONSUMR_PARTY_A"`
    PokeaCPI                              string `mapstructure:"POKEA_CREDIT_PARTY_IDENTIFIER"`
    PokeaBusinessExpressCheckoutShortCode string `mapstructure:"POKEA_BUSINESS_EXPRESS_CHECKOUT_SHORT_CODE"`
}

// ValidateAuth performs minimal validation necessary to authenticate against the M-Pesa API.
// It intentionally focuses on the values required to obtain an access token so that
// client construction can succeed for most use cases. Operation-specific validation
// should be performed by the respective service methods when needed.
func (c *Configs) ValidateAuth() error {
    if c.PokeaConsumerKey == "" || c.PokeaConsumerSecret == "" {
        return errors.New("missing POKEA_CONSUMER_KEY or POKEA_CONSUMER_SECRET")
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
    v.SetDefault("POKEA_TRANSACTION_TYPE", "CustomerPayBillOnline")

    // Read .env when available; ignore if missing
    if err := v.ReadInConfig(); err != nil {
        // If the file is missing, proceed with env vars only
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, err
        }
    }

    bindEnv(v)

    var configs Configs
    if err := v.Unmarshal(&configs); err != nil {
        return nil, err
    }

    // Backstop defaults when tags were not populated
    if configs.MpesaEnvironment == "" {
        configs.MpesaEnvironment = "sandbox"
    }
    if configs.PokeaTransactionType == "" {
        configs.PokeaTransactionType = "CustomerPayBillOnline"
    }

    if err := configs.ValidateAuth(); err != nil {
        return nil, err
    }
    return &configs, nil
}

// legacyEnvPrefix is the deprecated prefix for POKEA_* settings.
const legacyEnvPrefix = "DARAJA_"

// bindEnv binds every setting to its environment variable so OS env vars are
// read even without a .env file. Each POKEA_* setting also falls back to its
// deprecated DARAJA_* counterpart, from the OS environment or the .env file.
// Deprecated DARAJA_* names will be removed in the next major release.
func bindEnv(v *viper.Viper) {
    t := reflect.TypeOf(Configs{})
    for i := 0; i < t.NumField(); i++ {
        key := t.Field(i).Tag.Get("mapstructure")
        if !strings.HasPrefix(key, "POKEA_") {
            _ = v.BindEnv(key, key)
            continue
        }
        legacy := legacyEnvPrefix + strings.TrimPrefix(key, "POKEA_")
        _ = v.BindEnv(key, key, legacy)

        _, inEnv := os.LookupEnv(key)
        _, inLegacyEnv := os.LookupEnv(legacy)
        if !inEnv && !inLegacyEnv && !v.InConfig(key) && v.InConfig(legacy) {
            v.Set(key, v.Get(legacy))
        }
    }
}
