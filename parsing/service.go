package parsing

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"

	"github.com/alzedd/golb/settings"
	"github.com/alzedd/golb/taxonomy"

	"github.com/ericaro/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type settingsGetter interface {
	Get(name string) (v string)
}

type pathResolver interface {
	GetFileContent(file string) ([]byte, error)
	GetLayoutsFolder() string
	GetBlocksFolder() string
	GetFilesList(dir string) ([]string, error)
	GetContentSrcFolder() string
}

var settingsgetter = settings.Instance()

type service struct {
	store        PageGetter
	pathresolver pathResolver
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

	markDownData, err = s.pathresolver.GetFileContent(fn)

	if err == nil {
		p.MarkDownContent = string(markDownData)
		p.BlogName = settingsgetter.Get("BLOG_NAME")
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
		s.setPageLayout(settingsgetter.Get("DEFAULT_LAYOUT"), p)
	}
}

func (s *service) setPageLayout(l string, p *taxonomy.Page) {
	if len(l) == 0 {
		p.Layout = settingsgetter.Get("DEFAULT_LAYOUT")
	} else {
		p.Layout = l
	}
}

func (s *service) setPageFullHtml(p *taxonomy.Page) {
	buff := &bytes.Buffer{}
	layoutFolder := s.pathresolver.GetLayoutsFolder()
	layout := path.Join(layoutFolder, p.Layout+".html")

	blocksFolder := s.pathresolver.GetBlocksFolder()
	blocks, _ := s.pathresolver.GetFilesList(blocksFolder)

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
	opts := html.RendererOptions{
		Flags:          html.CommonFlags,
		RenderNodeHook: handleCheckboxes,
	}
	renderer := html.NewRenderer(opts)
	parsedHtml := string(markdown.ToHTML(([]byte)(p.MarkDownContent), parser, renderer))
	p.HtmlContent = template.HTML(parsedHtml)
}

func (s *service) GetStore() *taxonomy.PageRepository {
	return s.store.(*taxonomy.PageRepository)
}

func (s *service) ParseAll() {
	folderList, err := s.pathresolver.GetFilesList(s.pathresolver.GetContentSrcFolder())
	fmt.Printf("%+v\n", folderList)

	if err == nil {
		for _, folder := range folderList {
			var fileList []string
			var err error

			folderFullPath := path.Join(s.pathresolver.GetContentSrcFolder(), folder)
			folderInfo, err2 := os.Stat(folderFullPath)
			if err2 != nil {
				fmt.Println(folderFullPath)
				fmt.Println(err2.Error())
				os.Exit(1)
			}

			if err == nil {
				if folderInfo.IsDir() {
					fileList, err = s.pathresolver.GetFilesList(folderFullPath)
				} else {
					folderFullPath = s.pathresolver.GetContentSrcFolder()
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

func handleCheckboxes(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {

	if _, ok := node.(*ast.Text); ok {
		parent := node.GetParent()

		if _, ok := parent.(*ast.Paragraph); !ok {
			return ast.GoToNext, false
		}

		listItem := parent.GetParent()

		if _, isListItem := listItem.(*ast.ListItem); isListItem {
			text := node.AsLeaf().Literal

			switch {
			case bytes.HasPrefix(text, []byte("[ ] ")):
				text = append([]byte(`<input type="checkbox" disabled="">`), text[3:]...)
			case bytes.HasPrefix(text, []byte("[x] ")) || bytes.HasPrefix(text, []byte("[X] ")):
				text = append([]byte(`<input type="checkbox" checked="" disabled="">`), text[3:]...)
			}

			w.Write([]byte(text))
			return ast.GoToNext, true
		}
	} else if _, ok := node.(*ast.ListItem); ok {
		for _, child := range node.GetChildren() {
			if _, ok := child.(*ast.Paragraph); ok {
				text := child.GetChildren()[0].AsLeaf().Literal
				if bytes.HasPrefix(text, []byte("[ ] ")) || bytes.HasPrefix(text, []byte("[x] ")) || bytes.HasPrefix(text, []byte("[X] ")) {
					newText := ""
					if entering {
						newText = "<li class=\"task-list-item\">"
					} else {
						newText = "</li>"
					}
					w.Write([]byte(newText))
					return ast.GoToNext, true
				}
			}
		}

	}

	return ast.GoToNext, false
}
