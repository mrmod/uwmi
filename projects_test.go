package main

import (
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
)

func projectFactory() Project {
	return Project{Name: "project1", Description: "Testing project"}
}

func TestSave(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()
	if err := project.Save(ctx); err != nil {
		t.Error("Failed to save:", err)
	}
}

func TestOne(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()

	noProject := Project{Key: project.Key}
	if err := noProject.One(ctx); err != nil {
		t.Error("Expected an error, got none: ", err)
	}

	project.Save(ctx)

	foundProject := Project{Key: project.Key}
	err := foundProject.One(ctx)
	if err != nil {
		t.Error("Expected no errors, ", err)
	}
	if foundProject.Name != project.Name {
		t.Error("Failed to getOne", foundProject.Name)
	}
}

func TestAllByName(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()
	project.Save(ctx)
	time.Sleep(100 * time.Millisecond)
	projects, err := project.AllByName(ctx)
	if err != nil {
		t.Error("Expected no error, ", err)
	}
	if len(projects) != 1 {
		t.Error("Expected 1 projet")
	}
	if key := projects[0].Key; key != project.Key {
		t.Error("Expected the key", project.Key, "got", key)
	}
}

func TestAllProjects(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()
	project.Save(ctx)
	time.Sleep(100 * time.Millisecond)
	projects := AllProjects(ctx)
	if l := len(projects); l < 1 {
		t.Error("Expected to find at least one project")
	}
}

func TestTasks(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()
	project.Save(ctx)
	time.Sleep(100 * time.Millisecond)
	task := Task{Description: "Something", Project: &project}
	if err := task.Save(ctx); err != nil {
		t.Error("Failed to save", err)
	}
	time.Sleep(100 * time.Millisecond)
	err := project.AllTasks(ctx)
	if err != nil {
		t.Error("Expected no errors: ", err)
	}

	if len(project.Tasks) < 1 {
		t.Error("Expected 1 task")
	}
}

func TestSaveTask(t *testing.T) {
	ctx, _, _ := aetest.NewContext()
	project := projectFactory()
	project.Save(ctx)
	task := Task{Description: "Something", Project: &project}
	if err := task.Save(ctx); err != nil {
		t.Error("Expecte to save a task,", err)
	}
}
