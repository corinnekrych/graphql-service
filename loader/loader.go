package loader

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/dataloader"
)

type key string

const (
	UserLoaderKey     key = "user"
	WorkItemLoaderKey key = "wortkitem"
)

// Collection holds an internal lookup of initialized batch data load functions.
type BatchFnCollection struct {
	lookup map[key]dataloader.BatchFunc
}

// Attach creates new instances of dataloader.Loader and attaches the instances on the request context.
// We do this because the dataloader has an in-memory cache that is scoped to the request.
func (c BatchFnCollection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.lookup {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

// Initialize a lookup map of context keys to batch functions.
//
// When Attach is called on the Collection, the batch functions are used to create new dataloader
// instances. The instances are attached to the request context at the provided keys.
//
// The keys are then used to extract the dataloader instances from the request context.
func Initialize(client *client.Client) BatchFnCollection {
	return BatchFnCollection{
		lookup: map[key]dataloader.BatchFunc{
			UserLoaderKey: newUserLoader(client),
			//workItemLoaderKey:   newWorkItemLoader(client),
		},
	}
}
