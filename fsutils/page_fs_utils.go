package fsutils

import (
	"golb/settings"
	"io/ioutil"
	"os"
	"path"
)

func GetContentSrcFolder() string                { return path.Join(settings.Get("CONTENT_FOLDER")) }
func GetContentSrcFullPath(file string) string   { return path.Join(GetContentSrcFolder(), file) }
func GetFileContent(file string) ([]byte, error) { return ioutil.ReadFile(file) }

func GetContentDistFolder() string {
	p := path.Join(settings.Get("DIST_FOLDER"))
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, os.ModePerm)
	}
	return p
}

func GetContentDistFullPath(file string) string {
	p := path.Join(GetContentDistFolder(), file)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(p), os.ModePerm)
	}
	return p
}
