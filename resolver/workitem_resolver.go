package resolver

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/graphql-go"
)

// The WorkItemResolver is the entry point to retrieve work items.
type WorkItemResolver struct {
	wit client.WorkItem
}

func NewWorkItemsResolver(ctx context.Context, wits []client.WorkItem) (*[]*WorkItemResolver, error) {
	var resolvers = make([]*WorkItemResolver, 0, len(wits))
	for _, wit := range wits {
		resolvers = append(resolvers, &WorkItemResolver{wit: wit})
	}

	return &resolvers, nil
}

func (r WorkItemResolver) ID() graphql.ID {
	return graphql.ID(r.wit.ID.String())
}

// Description is a work item tracker's description.
func (r WorkItemResolver) Description() string {
	d := r.wit.Attributes["system.description"]
	if d == nil {
		return ""
	}
	if description, ok := d.(string); ok {
		return description
	}
	return ""
}

// Name is a work item tracker's title.
func (r WorkItemResolver) Name() string {
	t := r.wit.Attributes["system.title"]
	if t == nil {
		return ""
	}
	if title, ok := t.(string); ok {
		return title
	}
	return ""
}
