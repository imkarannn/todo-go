package main

import (

    "log"
    "net/http"
    "github.com/imkarannn/todo-go/handlers"
    
	
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)

    // Define API endpoints
    r.Post("/users", handlers.CreateUserEndpoint)
    r.Get("/users", handlers.GetUsersEndpoint)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}
