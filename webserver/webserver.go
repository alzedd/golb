package webserver

import (
	"golb/taxonomy"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *Router {

	pr := taxonomy.NewPageRepository()
	hc := newService(pr)
	assetsDir := pathresolver.GetAssetsPath()
	router := Router{}
	router.Router = mux.NewRouter().StrictSlash(true)
	router.Router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(assetsDir))))

	for _, route := range getRoutes(hc) {
		router.AddRoute(route)
	}
	return &router
}

func (router *Router) AddRoute(route Route) {
	router.Router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
}
