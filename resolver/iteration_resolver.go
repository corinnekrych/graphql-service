package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// The IterationResolver is the entry point to retrieve list of iteration for a given space.
type IterationResolver struct {
	spaceId   string
	iteration client.Iteration
	client    *client.Client
}

// NewIterationResolver creates an iteration resolver.
func NewIterationResolver(ctx context.Context, spaceId string, its []client.Iteration, client *client.Client) (*[]*IterationResolver, error) {
	var resolvers = make([]*IterationResolver, 0, len(its))
	for _, it := range its {
		resolvers = append(resolvers, &IterationResolver{spaceId: spaceId, iteration: it, client: client})
	}

	return &resolvers, nil
}

// ID is an unique iteration's id.
func (r IterationResolver) ID() graphql.ID {
	return graphql.ID(r.iteration.ID.String())
}

// Name is an iteration's name.
func (r IterationResolver) Name() string {
	if r.iteration.Attributes == nil {
		return ""
	}
	b := r.iteration.Attributes.Name
	if b == nil {
		return ""
	}
	return *b
}

// Description is an iteration's description.
func (r IterationResolver) Description() string {
	if r.iteration.Attributes == nil {
		return ""
	}
	d := r.iteration.Attributes.Description
	if d == nil {
		return ""
	}
	return *d
}

// State is an iteration's state: New, Closed...
func (r IterationResolver) State() string {
	if r.iteration.Attributes == nil {
		return ""
	}
	b := r.iteration.Attributes.State
	if b == nil {
		return ""
	}
	return *b
}

// WorkItems gets the list of work item trackers associated to an iteration.
func (r IterationResolver) WorkItems(ctx context.Context) (*[]*WorkItemResolver, error) {
	if r.iteration.ID == nil {
		return nil, nil
	}

	value := fmt.Sprintf("{\"$AND\":[{\"iteration\":{\"$EQ\":\"%s\"}}]}", r.iteration.ID.String())
	itJSON, err := r.client.ListWorkitems(ctx, "/api/search", nil, nil, nil, &value, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve WorkItems for IterationResolver")
	}
	defer itJSON.Body.Close()

	//dd, _ := ioutil.ReadAll(itJSON.Body)
	//fmt.Println(string(dd))
	var witData WorkItemsData
	err = json.NewDecoder(itJSON.Body).Decode(&witData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}

	resolver, err := NewWorkItemResolver(ctx, witData.Data, r.client)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create IterationResolver from filter")
	}
	return resolver, nil
}
