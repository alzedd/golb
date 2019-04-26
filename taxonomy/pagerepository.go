package taxonomy

import (
	"errors"
	"fmt"
)

type PageRepository struct {
	Pages map[string]*Page
}

func NewPageRepository() *PageRepository {
	Pages := map[string]*Page{}
	return &PageRepository{
		Pages,
	}
}

func (pr *PageRepository) Get(fname string) (*Page, error) {
	var err error = nil
	p, ok := pr.GetAll()[fname]

	if !ok {
		err = errors.New(fmt.Sprintf("no Pages found for key \"%s\"", fname))
		p = new(Page)
	}
	return p, err
}

func (p *PageRepository) GetAll() map[string]*Page {
	return p.Pages
}

func (pr *PageRepository) Add(p *Page) {
	pr.GetAll()[p.Filename] = p
}

func (pr *PageRepository) Delete(filename string) {
	delete(pr.GetAll(), filename)
}
