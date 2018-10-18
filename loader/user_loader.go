package loader

import (
	"context"
	"fmt"
	"github.com/corinnekrych/graphql-service/witapi/client"
	"github.com/graph-gophers/dataloader"
	"sync"
)

func newUserLoader(client *client.Client) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		result := []*dataloader.Result{}
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(len(keys))
		for _, id := range keys {
			go func(id dataloader.Key) {
				defer wg.Done()
				fmt.Println("::RESOLVING UserID: %v", id)
				path := fmt.Sprintf("/api/users/%s", id)
				userJSON, err := client.ShowUsers(ctx, path)
				if err != nil {
					mu.Lock()
					result = append(result, &dataloader.Result{
						Error: err,
					})
					mu.Unlock()
				}
				defer userJSON.Body.Close()
				user, err := client.DecodeUser(userJSON)
				if err != nil {
					mu.Lock()
					result = append(result, &dataloader.Result{
						Error: err,
					})
					mu.Unlock()
				}
				mu.Lock()
				result = append(result, &dataloader.Result{
					Data: *user,
				})
				mu.Unlock()
			}(id)
		}
		wg.Wait()
		return result
	}
}
