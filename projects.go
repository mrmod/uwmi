package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewProject(request *http.Request) (Project, error) {
	var project Project
	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)

	if err == nil {
		err = json.Unmarshal(body, &project)
		return project, err
	}

	return project, err
}
