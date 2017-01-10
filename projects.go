package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
)

// Project Is a container for tasks and documents
type Project struct {
	Key int64 `json:"key"`
	// Name of the project
	Name string `datastore:"name" json:"name"`
	// Description of the project
	Description string `datastore:"description,noindex" json:"description"`
	// Tasks list of tasks
	Tasks []Task `datastore:"-" json:"tasks"`
	// Developers attached to a project
	Developers []Developer `datastore:"-" json:"developers"`
	// Docs are specirication documents and stuff
	Docs    []Doc  `datastore:"-" json:"docs"`
	Created string `datastore:"created" json:"created"`
	Touched string `datastore:"touched" json:"touched"`
}

// NewProject Created from an HTTP request
func NewProject(request *http.Request) (Project, error) {
	var project Project
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)

	if err == nil {
		err = json.Unmarshal(body, &project)
		return project, err
	}
	// Set created and touched time to now
	project.Created = time.Now().UTC().Format(time.RFC3339)
	project.Touched = project.Created
	return project, err
}

// Save save a project to Datastore
func (self *Project) Save(ctx context.Context) error {
	projectKey, err := Save(ctx, ProjectKind, self, nil)
	if err != nil {
		return err
	}
	self.Key = projectKey.IntID()
	fmt.Printf("Created: %#v\n", self)
	return nil
}

// One project by Key
func (self *Project) One(ctx context.Context) error {
	key := datastore.NewKey(ctx, ProjectKind, "", self.Key, nil)
	return One(ctx, self, key)
}

// AllByName all projects matching this name
func (self *Project) AllByName(ctx context.Context) ([]Project, error) {
	log.Println("Finding project by name ", self.Name)
	var projects []Project
	keys, err := datastore.NewQuery(ProjectKind).Filter("name =", self.Name).GetAll(ctx, &projects)
	if err != nil {
		return projects, err
	}
	for i, key := range keys {
		projects[i].Key = key.IntID()
	}
	return projects, nil
}

func (self *Project) Delete(ctx context.Context) error {
	log.Println("Deleting ", self.Key)
	key := datastore.NewKey(ctx, ProjectKind, "", self.Key, nil)
	return datastore.Delete(ctx, key)
}

func (self *Project) AllTasks(ctx context.Context) ([]Task, error) {
	log.Println("Tasks ", self.Key)
	var tasks []Task
	keys, err := datastore.NewQuery(TaskKind).Filter("projectKey =", self.Key).GetAll(ctx, &tasks)
	if err != nil {
		return tasks, err
	}
	for i, key := range keys {
		tasks[i].Key = key.IntID()
	}
	self.Tasks = tasks
	return tasks, nil
}

// CreateTime
func (self Project) CreateTime() time.Time {
	return ParseTime(self.Created)
}

// TouchedTime
func (self Project) ModifiedTime() time.Time {
	return ParseTime(self.Touched)
}

func (self Project) DatastoreKey(ctx context.Context) *datastore.Key {
	return datastore.NewKey(ctx, ProjectKind, "", self.Key, nil)
}
