package resolver

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi"
	"github.com/pkg/errors"
)

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *witapi.Client
}

func NewRoot(client *witapi.Client) (*QueryResolver, error) {
	if client == nil {
		return nil, errors.New("cannot resolve witapi.Client")
	}

	return &QueryResolver{client: client}, nil
}

func (r QueryResolver) WorkItems(ctx context.Context) (*[]*WorkItemResolver, error) {
	// TODO use DataLoader
	wit, err := r.client.WorkItems(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve work items for QueryResolver")
	}
	resolver, err := NewWorkItemsResolver(ctx, wit)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create WorkItemsResolver")
	}
	return resolver, nil
}

type FilterQueryArgs struct {
	SpaceId string
}

func (r QueryResolver) Filter(ctx context.Context, args FilterQueryArgs) (*[]*WorkItemResolver, error) {
	// TODO use DataLoader
	wit, err := r.client.WorkItems(ctx, &args.SpaceId)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve filter for QueryResolver")
	}
	resolver, err := NewWorkItemsResolver(ctx, wit)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create WorkItemsResolver fro filter")
	}
	return resolver, nil
}
