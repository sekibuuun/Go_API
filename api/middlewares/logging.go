package middlewares

import (
	"log"
	"net/http"
)

// func(w http.ResponseWriter, req *http.Request) {} は無名関数
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI, req.Method)

		next.ServeHTTP(w, req)
	})
}
