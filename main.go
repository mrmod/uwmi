package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	appengine.Main()
	// See api.go for handlers
	router := mux.NewRouter()
	// Projects
	projectsRouter := router.PathPrefix("/api/projects").Subrouter()
	projectsRouter.HandleFunc("/", ProjectsHandler).Methods("GET", "POST")
	projectsRouter.HandleFunc("/{project}", ProjectHandler).Methods("GET", "PUT", "DELETE")
	// Tasks SubRoute
	projectsRouter.HandleFunc("/{project}/tasks", TasksHandler).Methods("GET", "POST")
	projectsRouter.HandleFunc("/{project}/tasks/{task}", TasksHandler).Methods("GET", "PUT", "DELETE")

	// Developers
	developersRouter := router.PathPrefix("/api/developers").Subrouter()
	developersRouter.HandleFunc("/", DevelopersHandler).Methods("GET", "POST")
	developersRouter.HandleFunc("/{developer}", DeveloperHandler).Methods("GET", "PUT", "DELETE")

	// Docs
	docsRouter := router.PathPrefix("/api/docs").Subrouter()
	docsRouter.HandleFunc("/", DocsHandler).Methods("GET", "POST")
	docsRouter.HandleFunc("/{doc}", DocHandler).Methods("GET", "PUT", "DELETE")

}
