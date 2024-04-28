package main

import (
    "context"
    "log"
    "net/http"
    "github.com/imkarannn/handlers"
	
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)

    // Define API endpoints
    r.Post("/users", CreateUserEndpoint)
    r.Get("/users", GetUsersEndpoint)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}
