package settings

import "github.com/spf13/viper"

type config struct {
	filetype string
	path     string
	name     string
}

var cfg *config = &config{
	filetype: "yaml",
	path:     ".",
	name:     ".golb",
}

var dfsettings map[string]interface{} = map[string]interface{}{
	"port":           3333,
	"blog_name":      "Blog",
	"theme":          "default",
	"default_layout": "default",
	"content_folder": "_content",
	"dist_folder":    "_release",
}

func init() {
	viper.SetConfigType(cfg.filetype)
	viper.AddConfigPath(cfg.path)
	viper.SetConfigName(cfg.name)

	for dfsettingkey, dfsettingval := range dfsettings {
		viper.BindEnv(dfsettingkey)
		viper.SetDefault(dfsettingkey, dfsettingval)
	}
}
