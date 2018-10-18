package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/pkg/errors"
)

type WorkItemMeta struct {
	TotalCount int `json:"totalCount"`
}
type WorkItemsData struct {
	Data []client.WorkItem `json:"data"`
	Meta WorkItemMeta      `json:"meta"`
}

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *client.Client
}

// NewRoot create a new root query resolver.
func NewRoot(client *client.Client) (*QueryResolver, error) {
	if client == nil {
		return nil, errors.New("cannot resolve witapi.Client")
	}

	return &QueryResolver{client: client}, nil
}

// FilterQueryArgs si the arguments to fetch work items.
type FilterQueryArgs struct {
	SpaceId string
}

// WorkItems fetches all the work items associated with a spaceId.
func (r QueryResolver) WorkItems(ctx context.Context, args FilterQueryArgs) (*[]*WorkItemResolver, error) {
	path := fmt.Sprintf("/api/spaces/%s/workitems", args.SpaceId)
	witJSON, err := r.client.ListWorkitems(ctx, path, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve workItems for QueryResolver")
	}

	defer witJSON.Body.Close()
	var witData WorkItemsData
	err = json.NewDecoder(witJSON.Body).Decode(&witData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}

	resolver, err := NewWorkItemResolver(ctx, witData.Data, r.client)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create WorkItemsResolver from filter")
	}
	return resolver, nil
}

type IterationsData struct {
	Data []client.Iteration `json:"data"`
}

// Iterations fetches all the iterations associated with a spaceId.
func (r QueryResolver) Iterations(ctx context.Context, args FilterQueryArgs) (*[]*IterationResolver, error) {
	path := fmt.Sprintf("/api/spaces/%s/iterations", args.SpaceId)
	witJSON, err := r.client.ListWorkitems(ctx, path, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve Iteration for QueryResolver")
	}

	defer witJSON.Body.Close()
	var itData IterationsData
	err = json.NewDecoder(witJSON.Body).Decode(&itData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}

	resolver, err := NewIterationResolver(ctx, args.SpaceId, itData.Data, r.client)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create IterationResolver from filter")
	}
	return resolver, nil
}

// WorkItems retirve the number of work item in the list for a given spaceId.
func (r QueryResolver) WorkItemsSummary(ctx context.Context, args FilterQueryArgs) (*WorkItemsSummaryResolver, error) {
	path := fmt.Sprintf("/api/spaces/%s/workitems", args.SpaceId)
	witJSON, err := r.client.ListWorkitems(ctx, path, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve workItemsSummary for QueryResolver")
	}
	defer witJSON.Body.Close()
	var witData WorkItemsData
	err = json.NewDecoder(witJSON.Body).Decode(&witData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}
	resolvedCount := 0
	inProgressCount := 0
	for _, wit := range witData.Data {
		state := wit.Attributes["system.state"]
		if state == nil {
			continue
		}
		st := state.(string)
		if st == "Resolved" {
			resolvedCount++
		}
		if st == "In Progress" {
			inProgressCount++
		}
	}
	resolver, err := NewWorkItemsSummayResolver(ctx, int32(witData.Meta.TotalCount), int32(resolvedCount), int32(inProgressCount))
	if err != nil {
		return nil, errors.Wrap(err, "cannot create WorkItemsSummaryResolver from filter")
	}
	return resolver, nil
}
