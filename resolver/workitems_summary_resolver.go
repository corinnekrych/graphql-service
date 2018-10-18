package resolver

import (
	"context"
)

// WorkItemsSummayResolver is the entry point to retrieve work items meta information.
type WorkItemsSummaryResolver struct {
	totalCount      int32
	resolvedCount   int32
	inProgressCount int32
}

// NewWorkItemsSummayResolver creates a work items list summary.
func NewWorkItemsSummayResolver(ctx context.Context, totalCount int32, resolvedCount int32, inProgressCount int32) (*WorkItemsSummaryResolver, error) {
	return &WorkItemsSummaryResolver{totalCount: totalCount, resolvedCount: resolvedCount, inProgressCount: inProgressCount}, nil
}

// TotalCount is the number of work items.
func (r WorkItemsSummaryResolver) TotalCount() int32 {
	return r.totalCount
}

// TotalCount is the number of work items.
func (r WorkItemsSummaryResolver) ResolvedCount() int32 {
	return r.resolvedCount
}

// TotalCount is the number of work items.
func (r WorkItemsSummaryResolver) InProgressCount() int32 {
	return r.inProgressCount
}
