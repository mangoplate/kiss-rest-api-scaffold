package server

import (
	"bitbucket.org/bearchit/books-api/handler"
	"github.com/bearchit/kiss"
	"github.com/gorilla/mux"
)

func routes(h *handler.Handler, k *kiss.Kiss) *mux.Router {
	return mux.NewRouter()
}
