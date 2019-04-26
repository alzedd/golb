package saving

import (
	"fmt"
	"golb/fsutils"
	"golb/parsing"
	"golb/taxonomy"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

type Pageparser interface {
	GetStore() *taxonomy.PageRepository
	ParseAll()
}

type Pagestorer interface {
	Get(fn string) (*taxonomy.Page, error)
	GetAll() map[string]*taxonomy.Page
	Add(*taxonomy.Page)
}

type service struct {
	parser Pageparser
}

func NewService(ps Pagestorer) *service {
	PageService := parsing.NewPageService(ps)
	return &service{parser: PageService}
}

func (s service) SaveAll() (err error) {
	s.parser.ParseAll()

	for _, p := range s.parser.GetStore().GetAll() {
		distPageFilename := strings.TrimSuffix(p.Filename, filepath.Ext(p.Filename))

		explPath := strings.Split(distPageFilename, string(os.PathSeparator))
		explPath = explPath[1:]
		distPageFilename = strings.Join(explPath, string(os.PathSeparator))

		newHtmlFile := fsutils.GetContentDistFullPath(distPageFilename + ".html")
		err = ioutil.WriteFile(newHtmlFile, []byte(p.GetFullHtml()), 0644)
		fmt.Printf("--:: Saving file :: %s\n", newHtmlFile)
	}

	return err
}

func (s service) SaveAssets() {
	srcAssetsDir := fsutils.GetAssetsPath()
	destAssetsDir := path.Join(fsutils.GetAssetsDestPath(), "static")
	copy.Copy(srcAssetsDir, destAssetsDir)
}
