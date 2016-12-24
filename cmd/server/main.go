package main

import (
	"fmt"
	"net/http"
	"os"

	"bitbucket.org/bearchit/books-api/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
	defer s.Close()

	s.Logger.Infof("Server started with port %d", s.Config.Server.Port)
	s.Logger.WithField("config", s.Config).Info()
	http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Server.Port), s)
}
