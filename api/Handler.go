package api

import (
	"net/http"
)

type devRoutes struct {
	Handler http.Handler
}

func (d *devRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	d.Handler.ServeHTTP(w, r)
}

func NewHandler(routes http.Handler, app *Application) http.Handler {
	apiRoutes := routes

	var newRoutes http.Handler
	serverRoutes := http.NewServeMux()
	newRoutes = serverRoutes
	serverRoutes.Handle("/api/", apiRoutes)
	if !app.IsDev() {
		serverRoutes.Handle("/", http.FileServer(http.Dir("./public")))
	} else {
		newRoutes = &devRoutes{serverRoutes}
	}

	return newRoutes
}
