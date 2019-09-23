package commands

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alzedd/golb/webserver"

	"github.com/jaschaephraim/lrserver"
	"github.com/rjeczalik/notify"
)

func Develop(s settingsGetter) {
	fmt.Printf("Starting development web server on http://localhost:%s\n", s.Get("PORT"))
	startLiveReload()
	router := webserver.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.Get("PORT")), router))
}

func startLiveReload() {
	lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
	go lr.ListenAndServe()

	c := make(chan notify.EventInfo, 1)
	if err := notify.Watch("./...", c, notify.Write); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			ei := <-c
			lr.Reload(ei.Path())
		}
	}()
}
