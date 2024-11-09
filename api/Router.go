package api

import (
	"encoding/json"
	"fmt"
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

type Response struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
	Content    any    `json:"response"`
}

func (r *Response) SetStatus(code int, text string) {
	r.Status = code
	r.StatusText = text
}

type Context struct {
	request  *http.Request
	response *Response
}

func (ctx *Context) Request() *http.Request {
	return ctx.request
}

type HandlerFn func(ctx *Context) any

type Router struct {
	staticContent  http.Handler
	apiGetContent  map[string]HandlerFn
	apiRouteRegexp regexp.Regexp
	middleware     http.HandlerFunc
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if router.middleware != nil {
		router.middleware(w, r)
	}

	if matcher := router.apiRouteRegexp.FindStringSubmatch(r.URL.Path); matcher != nil {
		ctx := &Context{
			request:  r,
			response: &Response{Status: http.StatusOK, StatusText: "OK"},
		}
		path := matcher[1]
		if path == "" {
			path = "/"
		}
		switch r.Method {
		case "GET":
			if route, ok := router.apiGetContent[path]; ok {
				ctx.response.Content = route(ctx)
				jsonB, err := json.Marshal(ctx.response)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "text/json")
				fmt.Fprint(w, string(jsonB))
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

func (r *Router) Get(path string, fn HandlerFn) {
	r.apiGetContent[path] = fn
}

func NewRouter(app *Application) *Router {
	router := Router{
		apiGetContent:  make(map[string]HandlerFn),
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
