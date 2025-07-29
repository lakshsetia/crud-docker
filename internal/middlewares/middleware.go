package middlewares

import "net/http"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authentication

		next.ServeHTTP(w, r)

		// logging
	})
}