package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {
	viper.AddConfigPath(".")
	viper.SetConfigName("optimizer")
	viper.SetConfigType("ini")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
