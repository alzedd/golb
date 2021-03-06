package parsing

import (
	"github.com/alzedd/golb/fsutils"
)

func NewPageService(pagestore PageGetter) *service {
	var pathresolver = fsutils.New()

	return &service{
		store:        pagestore,
		pathresolver: pathresolver,
	}
}
