package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const NoTrailingSlashRequired = true

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFoundHandler:", r.URL.Path)
}

//
func main() {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		log.Fatalln("Unable to get port")
	}
	// See api.go for handlers
	fmt.Println("Launching server on " + port)
	router := mux.NewRouter()
	router.StrictSlash(NoTrailingSlashRequired)
	router.NotFoundHandler = NotFoundHandler{}
	// Projects
	projectsRouter := router.PathPrefix("/api/projects").Subrouter()
	// projectsRouter.StrictSlash(NoTrailingSlashRequired)
	projectsRouter.StrictSlash(NoTrailingSlashRequired)
	projectsRouter.HandleFunc("/", ProjectsHandler).Methods("GET")
	projectsRouter.HandleFunc("/{projectKey}", ProjectHandler).Methods("GET")
	projectsRouter.HandleFunc("/", ProjectCreateHandler).Methods("POST")
	projectsRouter.HandleFunc("/{projectKey}", ProjectUpdateHandler).Methods("PUT")
	projectsRouter.HandleFunc("/{projectKey}", ProjectDeleteHandler).Methods("DELETE")
	// Tasks SubRoute
	projectsRouter.HandleFunc("/{projectKey}/tasks", TasksHandler).Methods("GET")
	projectsRouter.HandleFunc("/{projectKey}/tasks", TaskCreateHandler).Methods("POST")
	projectsRouter.HandleFunc("/{projectKey}/tasks/{taskKey}", TaskUpdateHandler).Methods("PUT")
	projectsRouter.HandleFunc("/{projectKey}/tasks/{taskKey}", TaskDeleteHandler).Methods("DELETE")

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
	fmt.Println("Registered all routes")
	// appengine.Main()
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("Error:", err)
	}

}
