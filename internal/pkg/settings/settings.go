package settings

import (
	"log"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func GetConfigPath() string     { return cfg.path }
func GetConfigName() string     { return cfg.name }
func GetConfigFileType() string { return cfg.filetype }

// ReadConfig: reads in config file and ENV variables if set.
func ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// load from environment variables or .env file
			viper.AutomaticEnv()
		}
	}
}

// writeConfig: writes an yml file with loaded settings
func WriteConfig(force bool) {
	filepath := path.Join(cfg.path, cfg.name+"."+strings.ToLower(cfg.filetype))

	if err := viper.WriteConfigAs(filepath); err != nil {
		log.Fatal(err.Error())
	}
}
