package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf sets up CSRF Token for all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad saves and loads the session state on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
