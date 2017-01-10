package main

import (
	"context"
	"fmt"

	"google.golang.org/appengine/datastore"
)

var (
	ErrorTooManyAncestors = fmt.Errorf("A maximum of one ancestor can be given")
)

func Save(ctx context.Context, kind string, object interface{}, parentKey *datastore.Key) (*datastore.Key, error) {
	newKey := datastore.NewIncompleteKey(ctx, kind, parentKey)
	return datastore.Put(ctx, newKey, object)
}

func One(ctx context.Context, object interface{}, key *datastore.Key) error {
	return datastore.Get(ctx, key, object)
}
