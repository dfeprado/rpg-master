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

func NewHandler(routes *http.ServeMux) http.Handler {
	apiRoutes := http.StripPrefix("/api", routes)
	if slices.Contains(os.Args, "--dev") {
		apiRoutes = &devRoutes{apiRoutes}
	}

	return apiRoutes
}
