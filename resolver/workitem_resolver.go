package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// The WorkItemResolver is the entry point to retrieve work items.
type WorkItemResolver struct {
	wit    client.WorkItem
	client *client.Client
}

func NewWorkItemResolver(ctx context.Context, wits []client.WorkItem, client *client.Client) (*[]*WorkItemResolver, error) {
	var resolvers = make([]*WorkItemResolver, 0, len(wits))
	for _, wit := range wits {
		resolvers = append(resolvers, &WorkItemResolver{wit: wit, client: client})
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

// Title is a work item tracker's title.
func (r WorkItemResolver) Title() string {
	t := r.wit.Attributes["system.title"]
	if t == nil {
		return ""
	}
	if title, ok := t.(string); ok {
		return title
	}
	return ""
}

// Name is a work item tracker's name.
func (r WorkItemResolver) Name() string {
	t := r.wit.Attributes["name"]
	if t == nil {
		return ""
	}
	if title, ok := t.(string); ok {
		return title
	}
	return ""
}

// State is a work item's state: New, Closed.
func (r WorkItemResolver) State() string {
	t := r.wit.Attributes["system.state"]
	if t == nil {
		return ""
	}
	if title, ok := t.(string); ok {
		return title
	}
	return ""
}

// Type is a work item's type: New, Closed.
func (r WorkItemResolver) Type() string {
	return r.wit.Type
}

type CommentData struct {
	Data []client.Comment `json:"data"`
}

func (r WorkItemResolver) Comments(ctx context.Context) (*[]*CommentResolver, error) {
	path := fmt.Sprintf("/api/workitems/%s/comments", r.wit.ID.String())
	witJSON, err := r.client.ListWorkItemComments(ctx, path, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve Comments for CommentResolver")
	}

	var commentData CommentData
	err = json.NewDecoder(witJSON.Body).Decode(&commentData)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to parse JSON: %v", err))
	}

	resolver, err := NewCommentResolver(ctx, commentData.Data)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create CommentResolver from filter")
	}
	return resolver, nil
}
