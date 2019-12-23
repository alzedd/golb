package fsutils

import (
	"os"
	"path"

	"github.com/alzedd/golb/internal/pkg/settings"
	"github.com/spf13/viper"
)

type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
}

func MkDirs(fs FileSystem) (err error) {
	settings.ReadConfig()

	if err = fs.MkdirAll(path.Join(settings.GetConfigPath(), viper.Get("content_folder").(string)), 0755); err == nil {
		err = fs.MkdirAll(path.Join(settings.GetConfigPath(), viper.Get("dist_folder").(string)), 0755)
	}

	return
}
