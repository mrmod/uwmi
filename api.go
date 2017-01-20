package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const (
	ProjectKind   = "project"
	TaskKind      = "task"
	DeveloperKind = "developer"
	DocKind       = "doc"

	ProjectResourceID   = "projectKey"
	TaskResourceID      = "taskKey"
	DeveloperResourceID = "developerKey"
	DocResourceID       = "docKey"

	dsEmptyStringID = ""
)

// Provide the ID for the request
func requestVar(request *http.Request, name string) string {
	return mux.Vars(request)[strings.ToLower(name)]
}

// RequetVarAsInt
func requestVarInt64(request *http.Request, name string) int64 {
	v := requestVar(request, name)
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("Error converting %s to int: %s\n", name, err)
	}
	return i
}

// ProjectsHandler Index handler
func ProjectsHandler(writer http.ResponseWriter, request *http.Request) {
	// var ctx context.Context
	// log.Println("Projects Index")
	// if v, ok := gctx.GetOk(request, "Context"); ok {
	// 	ctx = v.(context.Context)
	// } else {
	// 	ctx = appengine.NewContext(request)
	// }
	c := request.Context()
	log.Println("Context:", c)
	projects := AllProjects(appengine.NewContext(request))
	JSON(writer, projects)
}

// ProjectCreateHandler Create handler
func ProjectCreateHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ProjectCreateHandler")
	if project, err := NewProject(request); err != nil {
		ServerError(writer, err)
	} else {
		project.Save(appengine.NewContext(request))
		JSON(writer, project)
	}
}

// ProjectHandler Project handler
func ProjectHandler(writer http.ResponseWriter, request *http.Request) {
	key := requestVarInt64(request, ProjectResourceID)
	project := Project{Key: key}
	if err := project.One(appengine.NewContext(request)); err != nil {
		if err == datastore.ErrNoSuchEntity {
			NotFoundError(writer, ProjectKind, key, err)
		} else {
			ServerError(writer, err)
		}
	} else {
		JSON(writer, project)
	}
}

// ProjectUpdateHandler Update handler
func ProjectUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		BadRequest(writer, err)
		return
	}

	if err := project.Save(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
	} else {
		JSON(writer, project)
	}
}

// ProjectDeleteHandler Delete handler
func ProjectDeleteHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		BadRequest(writer, err)
		return
	}
	if err := project.Delete(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
	}
	writer.WriteHeader(http.StatusAccepted)
}

/*
Tasks Feature
*/
// TasksHandler Index handler
func TasksHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		ServerError(writer, err)
		return
	}
	if err := project.AllTasks(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
	} else {
		JSON(writer, project.Tasks)
		return
	}
}

// TaskHandler Task handler
func TaskHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		ServerError(writer, err)
		return
	}
	task, err := NewTask(request, &project)
	if err != nil {
		ServerError(writer, err)
		return
	}

	JSON(writer, task)
}

// TaskCreateHandler Create handler
func TaskCreateHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		ServerError(writer, err)
		return
	}

	task, err := NewTask(request, &project)
	if err != nil {
		ServerError(writer, err)
		return
	}

	if err = task.Save(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
		return
	}

	JSON(writer, task)
}

// TaskDeleteHandler Delete handler
func TaskDeleteHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		ServerError(writer, err)
		return
	}

	task, err := NewTask(request, &project)
	if err != nil {
		ServerError(writer, err)
		return
	}

	if err := task.Delete(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
		return
	}
	writer.WriteHeader(http.StatusAccepted)
}

// TaskUpdateHandler Update handler
func TaskUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	project, err := NewProject(request)
	if err != nil {
		ServerError(writer, err)
		return
	}

	task, err := NewTask(request, &project)
	if err != nil {
		ServerError(writer, err)
		return
	}

	if err := task.Save(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
		return
	}

	JSON(writer, task)
}

/*
Developers Feature
*/
// DevelopersHandler Index router
func DevelopersHandler(writer http.ResponseWriter, request *http.Request) {
	developers, err := AllDevelopers(appengine.NewContext(request))
	if err != nil {
		ServerError(writer, err)
		return
	}
	JSON(writer, developers)
}

// DeveloperCreateHandler create handler
func DeveloperCreateHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: Is a ServerError correct? Could this have been a malformed request?
	developer, err := NewDeveloper(request)
	if err != nil {
		ServerError(writer, err)
		return
	}

	if err := developer.Save(appengine.NewContext(request)); err != nil {
		ServerError(writer, err)
		return
	}

	JSON(writer, developer)
}

// DeveloperHandler Show handler
func DeveloperHandler(writer http.ResponseWriter, request *http.Request) {
	developer, err := NewDeveloper(request)
	if err != nil {
		ServerError(writer, err)
		return
	}

	JSON(writer, developer)
}

/*
Docs Feature
*/
// DocsHandler Index router
func DocsHandler(writer http.ResponseWriter, request *http.Request) {}
func DocHandler(writer http.ResponseWriter, request *http.Request)  {}
