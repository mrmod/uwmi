package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// ProjectDoc is a projects document
type ProjectDoc struct {
	Doc     *Doc     `datastore:"-" json:"doc"`
	Project *Project `datastore:"-" json:"project"`
}

// Doc is a thing with a URL
type Doc struct {
	Key int64 `json:"key"`
	// URL of the resource
	URL *url.URL `datastore:"url,noindex" json:"url"`
	// Description of what this is
	Description string `datastore:"description,noindex" json:"description"`
	Created     string `datastore:"created" json:"created"`
	// Touched The last time anything happened
	Touched string `datastore:"touched" json:"touched"`
}

// NewDoc Create a new document from an HTTP request
func NewDoc(request *http.Request) (Doc, error) {
	var doc Doc
	defer request.Body.Close()

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return doc, err
	} else if err := json.Unmarshal(body, &doc); err != nil {
		return doc, err
	}

	doc.Created = time.Now().UTC().Format(time.RFC3339)
	doc.Touched = doc.Created

	return doc, nil
}

// Save Save the document
func (self *Doc) Save(ctx context.Context) error {
	key, err := Save(ctx, DocKind, self, nil)
	self.Key = key.IntID()
	return err
}

// AddToProject Add a document to a project
func (self *Doc) AddToProject(ctx context.Context, project *Project) error {
	_, err := Save(ctx, DocKind, self, project.DatastoreKey(ctx))
	return err
}
