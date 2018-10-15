package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/pkg/errors"
)

type WorkItemsData struct {
	Data []client.WorkItem `json:"data"`
}

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *client.Client
}

func NewRoot(client *client.Client) (*QueryResolver, error) {
	if client == nil {
		return nil, errors.New("cannot resolve witapi.Client")
	}

	return &QueryResolver{client: client}, nil
}

type FilterQueryArgs struct {
	SpaceId string
}

func (r QueryResolver) WorkItems(ctx context.Context, args FilterQueryArgs) (*[]*WorkItemResolver, error) {
	// TODO use DataLoader
	//wit, err := r.client.WorkItems(ctx, &args.SpaceId)
	path := fmt.Sprintf("/api/spaces/%s/workitems", args.SpaceId)
	witJSON, err := r.client.ListWorkitems(ctx, path, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve workItems for QueryResolver")
	}

	var witData WorkItemsData
	err = json.NewDecoder(witJSON.Body).Decode(&witData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}

	resolver, err := NewWorkItemsResolver(ctx, witData.Data)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create WorkItemsResolver fro filter")
	}
	return resolver, nil
}
