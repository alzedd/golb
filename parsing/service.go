package parsing

import (
	"bytes"
	"errors"
	"fmt"
	"golb/fsutils"
	"golb/settings"
	"golb/taxonomy"
	"html/template"
	"os"
	"path"

	"github.com/ericaro/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type service struct {
	store PageGetter
}

type PageGetter interface {
	Get(fn string) (*taxonomy.Page, error)
	Add(p *taxonomy.Page)
}

func (s *service) Get(fn string, cached bool) (*taxonomy.Page, error) {
	fmt.Printf("--:: Parsing file :: %s\n", fn)
	p := new(taxonomy.Page)
	var err error = errors.New("NOT FOUND")

	if cached {
		p, err = s.store.Get(fn)
	}

	if err != nil {
		p, err = s.ParsePageFile(fn)
		s.store.Add(p)
	}

	return p, err
}

func (s *service) ParsePageFile(fn string) (p *taxonomy.Page, err error) {
	err = nil
	markDownData := []byte{}
	p = &taxonomy.Page{Filename: fn}

	markDownData, err = fsutils.GetFileContent(fn)

	if err == nil {
		p.MarkDownContent = string(markDownData)
		p.BlogName = settings.Get("BLOG_NAME")
		s.SetPageMetadata(markDownData, p)
		s.generateHtmlContent(p)
		s.setPageFullHtml(p)
	} else {
		err = errors.New(fmt.Sprintf("%s is not a valid MarkDown file or doesn't exist.", fn))
		fmt.Println(err.Error())
	}

	return
}

func (s *service) SetPageMetadata(PageContent []byte, p *taxonomy.Page) {
	frontmatter.Unmarshal(PageContent, p)
	if len(p.Layout) == 0 {
		s.setPageLayout(settings.Get("DEFAULT_LAYOUT"), p)
	}
}

func (s *service) setPageLayout(l string, p *taxonomy.Page) {
	if len(l) == 0 {
		p.Layout = settings.Get("DEFAULT_LAYOUT")
	} else {
		p.Layout = l
	}
}

func (s *service) setPageFullHtml(p *taxonomy.Page) {
	buff := &bytes.Buffer{}
	layoutFolder := fsutils.GetLayoutsFolder()
	layout := path.Join(layoutFolder, p.Layout+".html")

	blocksFolder := fsutils.GetBlocksFolder()
	blocks, _ := fsutils.GetFilesList(blocksFolder)

	for bIdx, b := range blocks {
		blocks[bIdx] = path.Join(blocksFolder, b)
	}

	blocks = append([]string{layout}, blocks...)

	tmpl := template.Must(template.ParseFiles(blocks...))
	tmpl.Execute(buff, p)

	p.FullHtml = buff.String()
}

func (s *service) generateHtmlContent(p *taxonomy.Page) {
	extensions := parser.CommonExtensions | parser.Mmark
	parser := parser.NewWithExtensions(extensions)
	parsedHtml := string(markdown.ToHTML(([]byte)(p.MarkDownContent), parser, nil))
	p.HtmlContent = template.HTML(parsedHtml)
}

func (s *service) GetStore() *taxonomy.PageRepository {
	return s.store.(*taxonomy.PageRepository)
}

func (s *service) ParseAll() {
	folderList, err := fsutils.GetFilesList(fsutils.GetContentSrcFolder())

	if err == nil {
		for _, folder := range folderList {
			var fileList []string
			var err error

			folderFullPath := path.Join(fsutils.GetContentSrcFolder(), folder)
			folderInfo, err2 := os.Stat(folderFullPath)
			if err2 != nil {
				fmt.Println(folderFullPath)
				fmt.Println(err2.Error())
				os.Exit(1)
			}

			if err == nil {
				if folderInfo.IsDir() {
					fileList, err = fsutils.GetFilesList(folderFullPath)
				} else {
					folderFullPath = fsutils.GetContentSrcFolder()
					fileList = []string{
						folder,
					}
				}

				for _, filename := range fileList {
					fileFullPath := path.Join(folderFullPath, filename)
					_, _ = s.Get(fileFullPath, true)
				}
			} else {
				fmt.Println(err.Error())
			}
		}
	}
}