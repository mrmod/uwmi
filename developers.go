package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
)

// DeveloperProject is a developer on a project
type DeveloperProject struct {
	Developer *Developer `datastore:"-" json:"developer"`
	Project   *Project   `datastore:"-" json:"project"`
}

// Developer is anyone that works on something
type Developer struct {
	Key   int64  `json:"key"`
	Name  string `datastore:"name" json:"name"`
	Email string `datastore:"email" json:"email"`
	// Projects the developer is a part of
	Projects []Project `datastore:"-" json:"projects"`
	// Tasks through projects
	Tasks []Task `datastore:"-" json:"tasks"`
	// FacebookID FacebookID
	FacebookID string `datastore:"facebookid" json:"facebookid"`
	// GoogleID GoogleID
	GoogleID string `datastore:"googleid" json:"googleid"`
	Created  string `datastore:"created" json:"created"`
	// Touched The last time anything happened
	Touched string `datastore:"touched" json:"touched"`
}

// NewDeveloper Create a new developer from an HTTP request
func NewDeveloper(request *http.Request) (Developer, error) {
	var developer Developer
	defer request.Body.Close()
	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return developer, err
	} else if err = json.Unmarshal(body, &developer); err != nil {
		return developer, err
	}

	developer.Projects = []Project{}
	developer.Tasks = []Task{}
	developer.Created = time.Now().UTC().Format(time.RFC3339)
	developer.Touched = developer.Created

	return developer, nil
}

func AllDevelopers(ctx context.Context) ([]Developer, error) {
	var developers []Developer
	keys, err := datastore.NewQuery(DeveloperKind).GetAll(ctx, developers)
	if err != nil {
		return developers, err
	}
	for i, key := range keys {
		developers[i].Key = key.IntID()
	}
	return developers, err
}

// Save Save the developer
func (self *Developer) Save(ctx context.Context) error {
	key, err := Save(ctx, DeveloperKind, self, nil)
	if err != nil {
		return err
	}
	self.Key = key.IntID()
	return err
}

// AddToProject Add a developer to a project
func (self *Developer) AddToProject(ctx context.Context, project Project) error {
	_, err := Save(ctx, DeveloperKind, self, project.DatastoreKey(ctx))
	return err
}

// AllProjects Fetch and set all projects for a developer
func (self *Developer) AllProjects(ctx context.Context, project Project) error {
	self.Projects = []Project{}
	keys, err := datastore.NewQuery(DeveloperKind).Ancestor(project.DatastoreKey(ctx)).GetAll(ctx, &self.Projects)
	if err != nil {
		return err
	}
	for i, key := range keys {
		self.Projects[i].Key = key.IntID()
	}
	return err
}

// AllTasks Fetch all tasks for a developer
func (self *Developer) AllTasks(ctx context.Context) {
	if len(self.Projects) == 0 {
		return
	}
	self.Tasks = []Task{}
	for _, project := range self.Projects {
		if err := project.AllTasks(ctx); err == nil {
			self.Tasks = append(self.Tasks, project.Tasks...)
		}
	}
}
