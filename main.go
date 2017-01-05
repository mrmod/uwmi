package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Provide the ID for the request
func requestVar(request *http.Request, name string) string {
	return mux.Vars(request)[strings.ToLower(name)]
}

/*
Projects Feature
*/
// ProjectsHandler Index handler
func ProjectsHandler(writer http.ResponseWriter, request *http.Request) {

}

// ProjectHandler Project handler
func ProjectHandler(writer http.ResponseWriter, request *http.Request) {
}

/*
Tasks Feature
*/
// TasksHandler Index router
func TasksHandler(writer http.ResponseWriter, request *http.Request) {

}

func TaskHandler(writer http.ResponseWriter, request *http.Request) {}

/*
Developers Feature
*/
// DevelopersHandler Index router
func DevelopersHandler(writer http.ResponseWriter, request *http.Request) {
}
func DeveloperHandler(writer http.ResponseWriter, request *http.Request) {}

/*
Docs Feature
*/
// DocsHandler Index router
func DocsHandler(writer http.ResponseWriter, request *http.Request) {}
func DocHandler(writer http.ResponseWriter, request *http.Request)  {}

func main() {
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
