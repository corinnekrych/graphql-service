package resolver

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/graphql-go"
)

// UserResolver is the entry point to retrieve user.
type UserResolver struct {
	user client.User
}

// NewUsersResolver creates a array of user resolver.
func NewUserResolver(ctx context.Context, users []client.User) (*[]*UserResolver, error) {
	var resolvers = make([]*UserResolver, 0, len(users))
	for _, user := range users {
		resolvers = append(resolvers, &UserResolver{user: user})
	}

	return &resolvers, nil
}

// ID of the user.
func (r UserResolver) ID() graphql.ID {
	if r.user.Data == nil {
		return ""
	}
	if r.user.Data.ID == nil {
		return ""
	}
	return graphql.ID(*r.user.Data.ID)
}

// Name is the full name of the user.
func (r UserResolver) Name() string {
	if r.user.Data == nil {
		return ""
	}
	if r.user.Data.Attributes == nil {
		return ""
	}
	if r.user.Data.Attributes.FullName == nil {
		return ""
	}
	return *r.user.Data.Attributes.FullName
}
