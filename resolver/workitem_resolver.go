package resolver

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/model"
	"github.com/graph-gophers/graphql-go"
)

// The WorkItemResolver is the entry point to retrieve work items.
type WorkItemResolver struct {
	wit model.WorkItem
}

func NewWorkItemsResolver(ctx context.Context, wits []model.WorkItem) (*[]*WorkItemResolver, error) {
	var resolvers = make([]*WorkItemResolver, 0, len(wits))
	for _, wit := range wits {
		resolvers = append(resolvers, &WorkItemResolver{wit: wit})
	}

	return &resolvers, nil
}

// ID resolves the film's unique identifier.
func (r WorkItemResolver) ID() graphql.ID {
	return graphql.ID(r.wit.Id)
}

// Description of a work item tracker
func (r WorkItemResolver) Description() string {
	return r.wit.WorkItemAttributes.Description
}

// Name of a work item tracker
func (r WorkItemResolver) Name() string {
	return r.wit.WorkItemAttributes.Title
}
