package resolver

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/graphql-go"
)

// CommentResolver is the entry point to retrieve work items.
type CommentResolver struct {
	comment *client.Comment
}

func NewCommentResolver(ctx context.Context, comments []client.Comment) (*[]*CommentResolver, error) {
	var resolvers = make([]*CommentResolver, 0, len(comments))
	for _, comment := range comments {
		resolvers = append(resolvers, &CommentResolver{comment: &comment})
	}

	return &resolvers, nil
}

func (r CommentResolver) ID() graphql.ID {
	return graphql.ID(r.comment.ID.String())
}

// Description is a work item tracker's description.
func (r CommentResolver) Body() string {
	if r.comment.Attributes == nil {
		return ""
	}
	b := r.comment.Attributes.Body
	if b == nil {
		return ""
	}
	return *b
}
