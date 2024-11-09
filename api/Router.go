package api

import (
	"net/http"
	"regexp"
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

type MiddlewareHandler struct {
	Handler      http.Handler
	MiddlewareFn http.HandlerFunc
}

func (m *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.MiddlewareFn(w, r)
	m.Handler.ServeHTTP(w, r)
}

type ApiRouter struct{}

type Router struct {
	staticContent  http.Handler
	apiGetContent  map[string]http.HandlerFunc
	apiRouteRegexp regexp.Regexp
	middleware     http.HandlerFunc
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if router.middleware != nil {
		router.middleware(w, r)
	}

	if matcher := router.apiRouteRegexp.FindStringSubmatch(r.URL.Path); matcher != nil {
		path := matcher[1]
		if path == "" {
			path = "/"
		}
		switch r.Method {
		case "GET":
			if route, ok := router.apiGetContent[path]; ok {
				route(w, r)
			} else {
				http.Error(w, "Not found", http.StatusNotFound)
			}
		default:
			http.Error(w, "Invalid method "+r.Method, http.StatusMethodNotAllowed)
		}
		return
	}

	if router.staticContent != nil {
		router.staticContent.ServeHTTP(w, r)
		return
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func (r *Router) Get(path string, fn http.HandlerFunc) {
	r.apiGetContent[path] = fn
}

func NewRouter(app *Application) *Router {
	router := Router{
		apiGetContent:  make(map[string]http.HandlerFunc),
		apiRouteRegexp: *regexp.MustCompile("^/api(/?.*)$"),
	}

	/* When the server starts with --dev (developing mode) arg, it adds the
	CORS middleware, so that the UI can fetch resources from the server.
	But, when this arg is not informed (production mode), the server must
	provide the static content at ./public path.*/
	if !app.IsDev() {
		router.staticContent = http.FileServer(http.Dir("./public"))
	} else {
		router.middleware = func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
		}
	}

	return &router
}
