package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"

	"github.com/gorilla/mux"
)

const (
	ProjectKind   = "project"
	TaskKind      = "task"
	DeveloperKind = "developer"
	DocKind       = "doc"

	dsEmptyStringID = ""
)

// Provide the ID for the request
func requestVar(request *http.Request, name string) string {
	return mux.Vars(request)[strings.ToLower(name)]
}

/*
Projects Feature
*/
// DatastoreEntity Is attached to Google datastore
type DatastoreEntity struct {
}

// ProjectsHandler Index handler
func ProjectsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectsHandler")
}

// ProjectCreateHandler Create handler
func ProjectCreateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectCreateHandler")
	project, err := NewProject(request)
	if err == nil {
		project.Save(appengine.NewContext(request))
		JSON(writer, project)
		return
	}
	ServerError(writer, err)
}

// ProjectHandler Project handler
func ProjectHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectHandler")
}

// ProjectUpdateHandler Update handler
func ProjectUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectUpdateHandler")
}

// ProjectDeleteHandler Delete handler
func ProjectDeleteHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectDeleteHandler")
}

/*
Tasks Feature
*/
// TasksHandler Index handler
func TasksHandler(writer http.ResponseWriter, request *http.Request) {}

// TaskHandler Task handler
func TaskHandler(writer http.ResponseWriter, request *http.Request) {}

// TaskCreateHandler Create handler
func TaskCreateHandler(writer http.ResponseWriter, request *http.Request) {}

// TaskDeleteHandler Delete handler
func TaskDeleteHandler(writer http.ResponseWriter, request *http.Request) {}

// TaskUpdateHandler Update handler
func TaskUpdateHandler(writer http.ResponseWriter, request *http.Request) {}

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
