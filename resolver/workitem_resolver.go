package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/corinnekrych/graphql-service/loader"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/dataloader"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

// The WorkItemResolver is the entry point to retrieve work items.
type WorkItemResolver struct {
	wit    client.WorkItem
	client *client.Client
}

// NewWorkItemResolver creates a work item resolver.
func NewWorkItemResolver(ctx context.Context, wits []client.WorkItem, client *client.Client) (*[]*WorkItemResolver, error) {
	var resolvers = make([]*WorkItemResolver, 0, len(wits))
	for _, wit := range wits {
		resolvers = append(resolvers, &WorkItemResolver{wit: wit, client: client})
	}

	return &resolvers, nil
}

// ID is the unique id of a work item tracker.
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

// CommentData is the root data Json type return REST call.
type CommentData struct {
	Data []client.Comment `json:"data"`
}

// Comments is the list of comments associated to a work item tracker.
func (r WorkItemResolver) Comments(ctx context.Context) (*[]*CommentResolver, error) {
	path := fmt.Sprintf("/api/workitems/%s/comments", r.wit.ID.String())
	witJSON, err := r.client.ListWorkItemComments(ctx, path, nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve Comments for CommentResolver")
	}
	defer witJSON.Body.Close()
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

// Assignees is the list of assignee associated to a work item tracker.
func (r WorkItemResolver) Assignees(ctx context.Context) (*[]*UserResolver, error) {

	if r.wit.Relationships == nil {
		return nil, nil
	}
	if r.wit.Relationships.Assignees == nil {
		return nil, nil
	}
	userIds := []string{}
	for _, v := range r.wit.Relationships.Assignees.Data {
		if v.ID == nil {
			continue
		}
		userIds = append(userIds, *v.ID)
	}
	users := []client.User{}

	// Load User with DataLoader
	//ldr := ctx.Value(loader.UserLoaderKey).(*dataloader.Loader)
	//keys := dataloader.NewKeysFromStrings(userIds)
	//for _, key := range keys {
	//	thunk := ldr.Load(ctx, key)
	//	data, err := thunk()
	//	if err != nil {
	//		return nil, err
	//	}
	//	if user, ok := data.(client.User); ok {
	//		users = append(users, user)
	//	}
	//}
	// end of Load User with DataLoader

	// Load User with DataLoader with Batch
	ldr := ctx.Value(loader.UserLoaderKey).(*dataloader.Loader)
	keys := dataloader.NewKeysFromStrings(userIds)
	thunks := ldr.LoadMany(ctx, keys)
	datas, errs := thunks()
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("%s", errs))
	}
	for _, data := range datas {
		if user, ok := data.(client.User); ok {
			users = append(users, user)
		}
	}
	// end of Load User with DataLoader

	// Load User without DataLoader
	//for _, id := range userIds {
	//	path := fmt.Sprintf("/api/users/%s", id)
	//	fmt.Println("::RESOLVING UserID %v", id)
	//	userJSON, err := r.client.ShowUsers(ctx, path)
	//	if err != nil {
	//		return nil, errors.Wrap(err, "cannot resolve ShowUsers for UserResolver")
	//	}
	//	defer userJSON.Body.Close()
	//	user, err := r.client.DecodeUser(userJSON)
	//	if err != nil {
	//		return nil, errors.Wrap(err, "cannot decode Users for UserResolver")
	//	}
	//	users = append(users, *user)
	//}
	// end of Load User without DataLoader
	resolver, err := NewUserResolver(ctx, users)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create CommentResolver from filter")
	}
	return resolver, nil
}
