package net

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"sync"
)

type Layout struct {
	Title         string
	PlayerAddress string
	Script        string
}

var layout Layout = Layout{
	Title:  "Master view - RPG-Master",
	Script: "/public/scripts/master/_layout.js",
}

func StartMasterServer(wg *sync.WaitGroup) {
	layout.PlayerAddress = "http://" + GetPlayerConnectionStruct().GetHostAndPort()
	address := "127.0.0.1:8080"
	fmt.Printf("You (master) can connect through http://%s\n", address)

	staticFiles := http.StripPrefix("/public/", http.FileServer(http.Dir("assets/public")))
	publicPath := regexp.MustCompile("^/public/?(.*)$")

	log.Fatal(http.ListenAndServe(address, &_HTTPHandler{
		func(w http.ResponseWriter, r *http.Request) {
			if m := publicPath.FindStringSubmatch(r.URL.Path); m != nil {
				staticFiles.ServeHTTP(w, r)
				return
			}

			// TODO move it to the global variables after dev
			var templates *template.Template = template.Must(template.ParseFiles(
				"./assets/web-pages/common/_layout.html",
			))

			templates.Execute(w, layout)
		},
	}))
	wg.Done()
}
