package loader

import (
	"context"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/dataloader"
)

func newUserLoader(client *client.Client) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		result := []*dataloader.Result{}
		for _, id := range keys {
			fmt.Println("::RESOLVING UserID %v", id)
			path := fmt.Sprintf("/api/users/%s", id)
			userJSON, err := client.ShowUsers(ctx, path)
			if err != nil {
				result = append(result, &dataloader.Result{
					Error: err,
				})
			}
			defer userJSON.Body.Close()
			user, err := client.DecodeUser(userJSON)
			if err != nil {
				result = append(result, &dataloader.Result{
					Error: err,
				})
			}
			result = append(result, &dataloader.Result{
				Data: *user,
			})
		}
		return result
	}
}
