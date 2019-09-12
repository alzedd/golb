package commands

import (
	"fmt"
	"golb/webserver"
	"log"
	"net/http"
)

func Develop(s settingsGetter) {
	fmt.Printf("Starting development web server on http://localhost:%s\n", s.Get("PORT"))
	router := webserver.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Get("PORT")), router))
}
