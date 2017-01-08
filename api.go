package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"

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
	Key int64 `json:"-"`
}

// Project Is a container for tasks and documents
type Project struct {
	DatastoreEntity
	// Name of the project
	Name string `datastore:"name" json:"name"`
	// Description of the project
	Description string `datastore:"description,noindex" json:"description"`
	// Tasks list of tasks
	Tasks []Task `datastore:"-" json:"tasks"`
	// Developers attached to a project
	Developers []Developer `datastore:"-" json:"developers"`
	// Docs are specirication documents and stuff
	Docs []Doc `datastore:"-" json:"docs"`
}

// Task Describes something to do. A task can only be given to one developer
type Task struct {
	DatastoreEntity
	// Description of the task
	Description string `datastore:"description,noindex" json:"description"`
	// Effort I dunno, maybe this is stupid
	Effort float32 `datastore:"effort" json:"effort"`
	// Priority Defaults to 0.5 of 1 (top) and 0.0 (lowest)
	Priority float32 `datastore:"priority" json:"priority"`
	// Started When it was started. Nil means it's unstarted
	Started *time.Time `datastore:"started" json:"started"`
	// Created When the task was created
	Created *time.Time `datastore:"created" json:"created"`
	// Touched The last time anything happened
	Touched *time.Time `datastore:"touched" json:"touched"`
}

// Developer is anyone that works on something
type Developer struct {
	DatastoreEntity
	// Projects the developer is a part of
	Projects []Project `datastore:"-" json:"projects"`
	// Tasks through projects
	Tasks []Task `datastore:"-" json:"tasks"`
	// FacebookID FacebookID
	FacebookID string `datastore:"facebookid" json:"facebookid"`
	// GoogleID GoogleID
	GoogleID string `datastore:"googleid" json:"googleid"`
}

// Doc is a thing with a URL
type Doc struct {
	DatastoreEntity
	// URL of the resource
	URL *url.URL `datastore:"url,noindex" json:"url"`
	// Description of what this is
	Description string `datastore:"description,noindex" json:"description"`
}

// Save save a project to Datastore
func (self *Project) Save() {
	context := context.Background()

	projectKey, err := datastore.Put(context,
		datastore.NewIncompleteKey(context, ProjectKind, nil),
		&self,
	)

	if err != nil {
		fmt.Println("Error Saving Project:", err)
		return
	}

	self.Key = projectKey.IntID()
	fmt.Printf("Created: %#v\n", self)
}

func (self Project) One() Project {
	var project Project
	context := context.Background()

	projectKey := datastore.NewKey(context, ProjectKind, dsEmptyStringID, self.Key, nil)
	if err := datastore.Get(context, projectKey, &project); err != nil {
		fmt.Println("Error getting Project", self.Key, ":", err)
	}

	return project
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
		project.Save()
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
