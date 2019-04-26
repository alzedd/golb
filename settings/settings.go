package settings

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var defaultSettings map[string]string = getDefaultSettings()

func Get(name string) (val string) {
	val = os.Getenv(name)
	if len(val) == 0 {
		val = defaultSettings[name]
	}
	return val
}

func getDefaultSettings() (dfSettings map[string]string) {
	dfSettings = map[string]string{
		"PORT":           "1234",
		"BLOG_FOLDER":    "blog",
		"BLOG_NAME":      "Blog",
		"THEME":          "default",
		"DEFAULT_LAYOUT": "default",
		"CONTENT_FOLDER": "_content",
		"DIST_FOLDER":    "dist",
	}
	return
}
