package configs

import (
	"fmt"
	"phihc116/go-task/backend/shared"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./webapi/configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&shared.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
