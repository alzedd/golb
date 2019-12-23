package settings

import (
	"log"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Settings int

func (s *Settings) GetConfigPath() string     { return cfg.path }
func (s *Settings) GetConfigName() string     { return cfg.name }
func (s *Settings) GetConfigFileType() string { return cfg.filetype }

// ReadConfig: reads in config file and ENV variables if set.
func (s *Settings) ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// load from environment variables or .env file
			viper.AutomaticEnv()
		}
	}
}

// writeConfig: writes an yml file with loaded settings
func (s *Settings) WriteConfig(force bool) {
	filepath := path.Join(cfg.path, cfg.name+"."+strings.ToLower(cfg.filetype))

	if err := viper.WriteConfigAs(filepath); err != nil {
		log.Fatal(err.Error())
	}
}
