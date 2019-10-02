package settings

import (
	"github.com/spf13/viper"
)

var dfsettings map[string]interface{} = map[string]interface{}{
	"port":           3333,
	"blog_name":      "Blog",
	"theme":          "default",
	"default_layout": "default",
	"content_folder": "_content",
	"dist_folder":    "_release",
}

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".golb")

	for dfsettingkey, dfsettingval := range dfsettings {
		viper.BindEnv(dfsettingkey)
		// make sure that envs vars have the same names as the yml/default values
		viper.SetDefault(dfsettingkey, dfsettingval)
	}
}
