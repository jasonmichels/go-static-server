package middleware

import (
	"log"
	"net/http"

	"github.com/newrelic/go-agent"
)

// NewRelicMiddleware New relic middleware that only adds new relic if you provided credentials
func NewRelicMiddleware(app newrelic.Application, pattern string, next http.Handler) (string, http.Handler) {

	if app != nil {
		log.Println("New relic is active")
		return newrelic.WrapHandle(app, "/", next)
	}

	return pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware Logging middleware to log each incoming request
func LoggingMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("{Method: \"%v\", URL: \"%v\", RequestURI: \"%v\"}\n", r.Method, r.URL.Path, r.RequestURI)
	})
}
