package api

import (
	"net/http"
	"os"
	"slices"
)

type devRoutes struct {
	Handler http.Handler
}

func (d *devRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	d.Handler.ServeHTTP(w, r)
}

func NewHandler(routes http.Handler) http.Handler {
	apiRoutes := routes
	isDev := slices.Contains(os.Args, "--dev")

	var newRoutes http.Handler
	serverRoutes := http.NewServeMux()
	newRoutes = serverRoutes
	serverRoutes.Handle("/api/", apiRoutes)
	if !isDev {
		serverRoutes.Handle("/", http.FileServer(http.Dir("./public")))
	} else {
		newRoutes = &devRoutes{serverRoutes}
	}

	return newRoutes
}
