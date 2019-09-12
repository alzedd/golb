package settings

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var s *settings

type settings struct {
	defaults map[string]string
}

func Instance() *settings {
	if s == nil {
		s = &settings{
			defaults: getDefaultSettings(),
		}
	}
	return s
}

func (s *settings) Get(name string) (v string) {
	v = os.Getenv(name)
	if len(v) == 0 {
		v = s.defaults[name]
	}
	return v
}

func getDefaultSettings() (dfSettings map[string]string) {
	dfSettings = map[string]string{
		"PORT":           "3333",
		"BLOG_NAME":      "Blog",
		"THEME":          "default",
		"DEFAULT_LAYOUT": "default",
		"CONTENT_FOLDER": "_content",
		"DIST_FOLDER":    "_release",
	}
	return
}
