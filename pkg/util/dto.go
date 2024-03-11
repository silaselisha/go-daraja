package util

import "github.com/spf13/viper"

type Configs struct {
	DarajaEnvironment string `mapstructure:"DARAJA_ENVIRONMENT"`
	DarajaConsumerKey string `mapstructure:"DARAJA_CONSUMER_KEY"`
	DarajaConsumerSecret string `mapstructure:"DARAJA_CONSUMER_SECRET"`
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