package webserver

import (
	"golb/fsutils"
	"golb/parsing"
	"golb/taxonomy"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/afero"
)

type Pageparser interface {
	Get(fn string, cached bool) (*taxonomy.Page, error)
}

type htmlHolder interface {
	GetFullHtml() string
}

type Pagestorer interface {
	Get(fn string) (*taxonomy.Page, error)
	Add(*taxonomy.Page)
}

type service struct {
	parser Pageparser
}

func newService(ps Pagestorer) *service {
	PageService := parsing.NewPageService(ps)
	return &service{PageService}
}

func (s *service) PreviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PageFolder := vars["folder"]
	PageFileName := strings.TrimSuffix(vars["slug"], filepath.Ext(vars["slug"]))

	PageFileName = generatePageFilename(PageFolder, PageFileName)
	Page, err := s.parser.Get(fsutils.GetContentSrcFullPath(PageFileName)+".md", false)

	if err == nil {
		tmpl, err := s.getPrintablePage(Page)
		if err == nil {
			tmpl.Execute(w, Page)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
}

func (s *service) getPrintablePage(p htmlHolder) (*template.Template, error) {
	return template.New("developmentPage").Parse(p.GetFullHtml())
}

func generatePageFilename(folder string, file string) string {
	fs := afero.NewOsFs()

	if len(folder) == 0 && len(file) == 0 {
		file = "index"
	} else {
		if len(file) == 0 {
			file = strings.TrimSuffix(folder, path.Ext(folder))
			checkfile := fsutils.GetContentSrcFullPath(file)
			if _, err := fs.Stat(checkfile + ".md"); os.IsNotExist(err) {
				file = path.Join(file, "index")
			}
		} else {
			file = path.Join(folder, file)
		}
	}

	return file
}
