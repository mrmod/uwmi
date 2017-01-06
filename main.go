package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const NoTrailingSlashRequired = true

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundHandler:", r.URL.Path)
}

func main() {
	// appengine.Main()
	// See api.go for handlers
	fmt.Println("Launching server")
	router := mux.NewRouter()
	router.StrictSlash(NoTrailingSlashRequired)
	router.NotFoundHandler = NotFoundHandler{}
	// Projects
	projectsRouter := router.PathPrefix("/api/projects").Subrouter()
	// projectsRouter.StrictSlash(NoTrailingSlashRequired)
	projectsRouter.HandleFunc("/", ProjectsHandler).Methods("GET", "POST")
	projectsRouter.HandleFunc("/{project}", ProjectHandler).Methods("GET", "PUT", "DELETE")
	// Tasks SubRoute
	projectsRouter.HandleFunc("/{project}/tasks", TasksHandler).Methods("GET", "POST")
	projectsRouter.HandleFunc("/{project}/tasks/{task}", TasksHandler).Methods("GET", "PUT", "DELETE")

	// Developers
	developersRouter := router.PathPrefix("/api/developers").Subrouter()
	developersRouter.StrictSlash(NoTrailingSlashRequired)
	developersRouter.HandleFunc("/", DevelopersHandler).Methods("GET", "POST")
	developersRouter.HandleFunc("/{developer}", DeveloperHandler).Methods("GET", "PUT", "DELETE")

	// Docs
	docsRouter := router.PathPrefix("/api/docs").Subrouter()
	docsRouter.StrictSlash(NoTrailingSlashRequired)
	docsRouter.HandleFunc("/", DocsHandler).Methods("GET", "POST")
	docsRouter.HandleFunc("/{doc}", DocHandler).Methods("GET", "PUT", "DELETE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error:", err)
	}
}
