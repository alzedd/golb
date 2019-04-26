package webserver

import (
	"net/http"
)

type handlerContainer interface {
	PreviewHandler(w http.ResponseWriter, r *http.Request)
}

func newHandlerContainer(store Pagestorer) handlerContainer {
	return newService(store)
}

func getRoutes(s handlerContainer) Routes {

	var routes = Routes{
		{
			Name:        "BlogIndex",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: s.PreviewHandler,
		},
		{
			Name:        "CategoryIndex",
			Method:      "GET",
			Pattern:     "/{folder}/",
			HandlerFunc: s.PreviewHandler,
		},
		{
			Name:        "PagePreview",
			Method:      "GET",
			Pattern:     "/{folder}/{slug}",
			HandlerFunc: s.PreviewHandler,
		},
	}

	return routes
}
