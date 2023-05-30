package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-remotekeyvalue/examples/api/exampleserver"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	var apiKeyNamePointer = flag.String("keyname", "API_KEY", "Name of the API Key Header")
	var apiKeyPointer = flag.String("key", "12345", "Allowed value of the API_KEY header")
	flag.Parse()

	pairs := exampleserver.GetDummyData()
	store := exampleserver.NewStore(pairs[:])

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Use(AuthorizeApiKey(DerefString(apiKeyNamePointer), DerefString(apiKeyPointer)))

	exampleserver.HandlerFromMux(store, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Printf("Starting Server at %s...", s.Addr)

	var routes = make([]string, 0)

	for _, v := range r.Routes() {
		routes = append(routes, v.Pattern)
	}

	log.Printf("Supported Routes: [ %s ]", strings.Join(routes, ", "))

	var keys = make([]string, 0)

	for _, v := range pairs[:] {
		keys = append(keys, v.Key)
	}

	log.Printf("Available Keys: [ %s ]", strings.Join(keys, ", "))

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

func AuthorizeApiKey(headerName string, apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var suppliedKey = r.Header.Get(headerName)
			if suppliedKey == "" || suppliedKey != apiKey {
				err := exampleserver.Error{
					Code:    http.StatusUnauthorized,
					Message: fmt.Sprintf("Supplied API Key header %s, is either missied or has an invalid key.", headerName),
				}
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func DerefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
