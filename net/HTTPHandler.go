package net

import "net/http"

type _HTTPHandler struct {
	callback func(http.ResponseWriter, *http.Request)
}

func (h *_HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.callback(w, r)
}