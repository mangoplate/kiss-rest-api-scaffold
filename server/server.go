package server

import (
	"bitbucket.org/bearchit/books-api/domain"
	"bitbucket.org/bearchit/books-api/handler"
	"bitbucket.org/bearchit/books-api/resource"
	"github.com/bearchit/kiss"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

type server struct {
	*negroni.Negroni
	*resource.Bundle
}

var (
	authWhitelist = []string{
		"/health_check",
	}
)

func NewServer() (*server, error) {
	rb, err := resource.NewBundle()
	if err != nil {
		return nil, err
	}

	requestLogger := negronilogrus.NewMiddlewareFromLogger(rb.Logger.Logger, "")
	requestLogger.ExcludeURL("/health_check")

	authenticator := kiss.NewNegroniMiddleware(domain.NewAuthenticator(rb, authWhitelist))

	n := negroni.New(
		negroni.NewRecovery(),
		requestLogger,
		authenticator,
	)

	k := kiss.New()
	h := handler.Handler{Bundle: rb}
	r := routes(&h, k)
	n.UseHandler(r)

	return &server{Negroni: n, Bundle: rb}, nil
}
