package commands

import (
	"fmt"
	"golb/settings"
	"golb/webserver"
	"log"
	"net/http"
)

func CommandDevelop() {
	fmt.Printf("Starting development web server on http://localhost:%s\n", settings.Get("PORT"))
	router := webserver.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", settings.Get("PORT")), router))
}
