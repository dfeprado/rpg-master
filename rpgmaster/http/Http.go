package http

import (
	"net/http"
	"regexp"
)

/*
Type that handles static content serving.

Use the function NewPublicHandler() to create one.
*/
type StaticHandler struct {
	URLPattern regexp.Regexp
	Handler    http.Handler
}

/*
Creates a new PublicHandler that will serve clients with static contents

httpPath must be a fqn path, ie, start with "/"" and end with "/", eg /ui/.

fsPath is the path on the file system that contains the desired static content.
*/
func NewPublicHandler(httpPath, fsPath string) StaticHandler {
	httpPath += "static/"
	return StaticHandler{
		URLPattern: *regexp.MustCompile("^" + httpPath + "?.*$"),
		Handler: http.StripPrefix(
			httpPath,
			http.FileServer(
				http.Dir(fsPath),
			),
		),
	}
}

/*
The HTTP Handler
*/
type Handler struct {
	Callback      func(http.ResponseWriter, *http.Request)
	PublicHandler *StaticHandler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.handlePublicHTTP(w, r) {
		return
	}
	h.Callback(w, r)
}

func (h *Handler) handlePublicHTTP(w http.ResponseWriter, r *http.Request) bool {
	if h.PublicHandler == nil {
		return false
	}

	if !h.PublicHandler.URLPattern.MatchString(r.URL.Path) {
		return false
	}

	h.PublicHandler.Handler.ServeHTTP(w, r)
	return true
}
