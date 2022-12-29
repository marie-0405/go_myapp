package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/marie-0405/go_myapp/pkg/config"
	"github.com/marie-0405/go_myapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
