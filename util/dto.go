package util

import "github.com/spf13/viper"

type Configs struct {
	DarajaEnvironment       string `mapstructure:"DARAJA_ENVIRONMENT"`
	DarajaConsumerKey       string `mapstructure:"DARAJA_CONSUMER_KEY"`
	DarajaConsumerSecret    string `mapstructure:"DARAJA_CONSUMER_SECRET"`
	DarajaBusinessShortCode string `mapstructure:"DARAJA_BUSINESS_SHORT_CODE"`
	DarajaPassKey           string `mapstructure:"DARAJA_PASS_KEY"`
	DarajaTransactionType   string `mapstructure:"DARAJA_TRANSACTION_TYPE"`
	DarajaCallBackURL       string `mapstructure:"DARAJA_CALL_BACK_URL"`
	DarajaPartyB            string `mapstructure:"DARAJA_PARTY_B"`
	DarajaTimestamp         string `mapstructure:"DARAJA_TIMESTAMP"`
	DarajaAccountRef        string `mapstructure:"DARAJA_ACCOUNT_REF"`
}

func LoadConfigs(path string) (envs *Configs, err error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&envs)
	if err != nil {
		return
	}
	return
}
