package witapi

import (
	"context"
	"github.com/corinnekrych/graphql-service/witapi/model"
	"github.com/pkg/errors"
)

type WorkItemData struct {
	Data []model.WorkItem `json:"data"`
}

func (c *Client) WorkItems(ctx context.Context, args *string) ([]model.WorkItem, error) {
	witList := []model.WorkItem{}
	var witData WorkItemData
	// TODO check WIT api for a default namespace
	endpoint := "/spaces/e8864cfe-f65a-4351-85a4-3a585d801b45/workitems"
	if args != nil {
		endpoint = "/spaces/" + *args + "/workitems"
	}
	err := c.Get(ctx, endpoint, &witData)
	if err != nil {
		return witList, errors.Wrap(err, "cannot fetch work items from wit api")
	}
	return witData.Data, nil
}
