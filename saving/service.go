package saving

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alzedd/golb/fsutils"
	"github.com/alzedd/golb/parsing"
	"github.com/alzedd/golb/settings"
	"github.com/alzedd/golb/taxonomy"

	"github.com/otiai10/copy"
)

type pathResolver interface {
	GetFileContent(file string) ([]byte, error)
	GetLayoutsFolder() string
	GetBlocksFolder() string
	GetFilesList(dir string) ([]string, error)
	GetContentSrcFolder() string
}

var s = settings.Instance()
var pathresolver = fsutils.New()

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

		newHtmlFile := pathresolver.GetContentDistFullPath(distPageFilename + ".html")
		err = ioutil.WriteFile(newHtmlFile, []byte(p.GetFullHtml()), 0644)
		fmt.Printf("--:: Saving file :: %s\n", newHtmlFile)
	}

	return err
}

func (s service) SaveAssets() {
	srcAssetsDir := pathresolver.GetAssetsPath()
	destAssetsDir := path.Join(pathresolver.GetAssetsDestPath(), "static")
	copy.Copy(srcAssetsDir, destAssetsDir)
}
