package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// TODO.Make conf file optional, default to curr values and ask for file if no match
func Load() {
	viper.AddConfigPath(".")
	viper.SetConfigName("optimizer")
	viper.SetConfigType("ini")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
