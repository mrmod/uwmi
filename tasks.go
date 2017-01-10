package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
)

// Task Describes something to do. A task can only be given to one developer
type Task struct {
	Key     int64    `json:"key"`
	Project *Project `datastore:"-" json:"project"`
	// Description of the task
	Description string `datastore:"description,noindex" json:"description"`
	// Effort I dunno, maybe this is stupid
	Effort float32 `datastore:"effort" json:"effort"`
	// Priority Defaults to 0.5 of 1 (top) and 0.0 (lowest)
	Priority float32 `datastore:"priority" json:"priority"`
	// Started When it was started. Nil means it's unstarted
	Started string `datastore:"started" json:"started"`
	// Created When the task was created
	Created string `datastore:"created" json:"created"`
	// Touched The last time anything happened
	Touched string `datastore:"touched" json:"touched"`
}

func NewTask(request *http.Request, project *Project) (Task, error) {
	var task Task
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		err = json.Unmarshal(body, &task)
		return task, err
	}

	task.Project = project
	task.Created = time.Now().UTC().Format(time.RFC3339)
	task.Touched = task.Created
	return task, err
}

func (self *Task) Save(ctx context.Context) error {
	key, err := Save(ctx, TaskKind, self, self.Project.DatastoreKey(ctx))
	if err != nil {
		return err
	}
	self.Key = key.IntID()
	log.Println("Saved ", key)
	return err
}

func (self *Task) One(ctx context.Context) error {
	return One(ctx, self, self.DatastoreKey(ctx))
}

func (self Task) DatastoreKey(ctx context.Context) *datastore.Key {
	return datastore.NewKey(ctx, TaskKind, "", self.Key, self.Project.DatastoreKey(ctx))
}
