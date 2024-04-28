package main

import (
    "context"
    "log"
    "net/http"
	"D:/Repos/GO/TodoProject/handlers"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
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
