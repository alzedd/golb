package fsutils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
}

type configReader interface {
	GetConfigPath() string
}

func MkDirs(fs FileSystem, settings configReader) (err error) {
	if err = fs.MkdirAll(path.Join(settings.GetConfigPath(), viper.Get("content_folder").(string)), 0755); err == nil {
		err = fs.MkdirAll(path.Join(settings.GetConfigPath(), viper.Get("dist_folder").(string)), 0755)
	}

	return
}
