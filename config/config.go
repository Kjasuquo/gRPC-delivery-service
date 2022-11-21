package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GrpcPort      string `mapstructure:"GRPC_PORT" json:"GRPC_PORT"`
	ConsulAddress string `mapstructure:"CONSUL_ADDRESS" json:"consulAddress"`
}

func ReadConfigs(path string) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.AutomaticEnv()

		} else {
			fmt.Printf("cannot read config: %v", err)
		}
	}

	config, err := VaultSecrets(viper.GetString("VAULT_ADDR"), viper.GetString("VAULT_AUTH_TOKEN"), viper.GetString("VAULT_SECRET_PATH"))

	if err != nil {
		fmt.Println(err, "couldn't load secrets")
	}

	configs := &Config{
		GrpcPort: config.GrpcPort,
	}

	return configs

}
