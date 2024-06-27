package middleware

import (
	"fmt"
	"net/http"
)

func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		fmt.Println("Executing middlewareOne again")
	})
}
