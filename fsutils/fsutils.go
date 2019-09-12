package fsutils

import (
	"io/ioutil"
	"path"
)

type fsutils struct{}

func New() *fsutils {
	return &fsutils{}
}

func (fsu *fsutils) GetFilesList(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	var fileList []string
	if err == nil {
		for _, file := range files {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList, err
}

func (fsu *fsutils) GetAssetsPath() string {
	return path.Join("_themes", settingsgetter.Get("THEME"), "assets")
}

func (fsu *fsutils) GetLayoutsFolder() string {
	return path.Join("_themes", settingsgetter.Get("THEME"), "layouts")
}

func (fsu *fsutils) GetBlocksFolder() string {
	return path.Join("_themes", settingsgetter.Get("THEME"), "blocks")
}

func (fsu *fsutils) GetAssetsDestPath() string {
	return path.Join(settingsgetter.Get("DIST_FOLDER"))
}
