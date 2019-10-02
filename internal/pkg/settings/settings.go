package settings

import (
	"fmt"

	"github.com/spf13/viper"
)

// ReadConfig reads in config file and ENV variables if set.
func ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.AutomaticEnv()
		}
	}
}

// writeConfig writes a yml file with default settings
func WriteConfig() {
	if err := viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(err.Error())
		}
	}
}
