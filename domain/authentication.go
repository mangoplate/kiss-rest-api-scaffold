package domain

import (
	"net/http"
	"regexp"
	"strings"

	"bitbucket.org/bearchit/books-api/resource"
	"github.com/bearchit/kiss"
)

type whitelist []*regexp.Regexp

type authenticator struct {
	bundle    *resource.Bundle
	whitelist whitelist
}

func NewAuthenticator(bundle *resource.Bundle, whitelist []string) *authenticator {
	a := authenticator{
		bundle: bundle,
	}

	for _, path := range whitelist {
		a.whitelist = append(a.whitelist, regexp.MustCompile(path))
	}

	return &a
}

func (a *authenticator) ServeHTTP(c *kiss.Context, next http.HandlerFunc) {
	for _, w := range a.whitelist {
		if w.MatchString(c.Request.RequestURI) {
			next(c.ResponseWriter, c.Request)
			return
		}
	}

	token := c.Request.Header.Get("Authorization")
	if strings.TrimSpace(token) == "" {
		c.Unauthorized(nil)
		return
	}

	// Authentication codes here...

	next(c.ResponseWriter, c.Request)
}
