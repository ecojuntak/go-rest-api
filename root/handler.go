package root

import (
	"net/http"
)

type rootHandler struct{}

type RootHandler interface {
	Root(w http.ResponseWriter, r *http.Request)
}

func NewRootHandler() RootHandler {
	return &rootHandler{}
}

func (h rootHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}
