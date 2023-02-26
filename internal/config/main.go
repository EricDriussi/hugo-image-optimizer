package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load(cfgFile string) {
	setDefaults()
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/")
	viper.SetConfigName("optimizer")
	viper.SetConfigType("ini")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file was found, sticking with default values")
	}
}

func setDefaults() {
	viper.SetDefault("dirs.posts", "content/posts/")
	viper.SetDefault("dirs.images", "static/images/")
	viper.SetDefault("dirs.images_exclude", []string{"whoami", "donation"})
	viper.SetDefault("compression.quality", 50)
}
