package fsutils

import (
	"io/ioutil"
	"os"
	"path"
)

func (fsu *fsutils) GetContentSrcFolder() string {
	return path.Join(settingsgetter.Get("CONTENT_FOLDER"))
}

func (fsu *fsutils) GetContentSrcFullPath(file string) string {
	return path.Join(fsu.GetContentSrcFolder(), file)
}

func (fsu *fsutils) GetFileContent(file string) ([]byte, error) { return ioutil.ReadFile(file) }

func (fsu *fsutils) GetContentDistFolder() string {
	p := path.Join(settingsgetter.Get("DIST_FOLDER"))
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, os.ModePerm)
	}
	return p
}

func (fsu *fsutils) GetContentDistFullPath(file string) string {
	p := path.Join(fsu.GetContentDistFolder(), file)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(p), os.ModePerm)
	}
	return p
}
