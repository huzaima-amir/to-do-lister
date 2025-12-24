package middleware

import (
	"net/http"
	"github.com/go-chi/cors"
)

func Cors() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
