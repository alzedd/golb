package fsutils

import (
	"golb/settings"
	"io/ioutil"
	"path"
)

func GetFilesList(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	var fileList []string
	if err == nil {
		for _, file := range files {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList, err
}

func GetAssetsPath() string {
	return path.Join("_themes", settings.Get("THEME"), "assets")
}

func GetLayoutsFolder() string {
	return path.Join("_themes", settings.Get("THEME"), "layouts")
}

func GetBlocksFolder() string {
	return path.Join("_themes", settings.Get("THEME"), "blocks")
}

func GetAssetsDestPath() string {
	return path.Join(settings.Get("DIST_FOLDER"))
}
