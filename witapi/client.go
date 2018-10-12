package witapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient creates a http client.
func NewClient(c *http.Client, baseURL string) *Client {
	if c == nil {
		c = http.DefaultClient
	}

	return &Client{baseURL: baseURL, client: c}
}

// Get handles json http GET requests.
func (c *Client) Get(ctx context.Context, endpoint string, object interface{}) error {
	req, err := http.NewRequest("GET", c.baseURL+endpoint, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("cannot create an http request for %v", c.baseURL))
	}
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("cannot call http.Get for %v", c.baseURL))
	}
	defer resp.Body.Close()
	if object != nil {
		if err = json.NewDecoder(resp.Body).Decode(object); err != nil {
			return errors.Wrap(err, fmt.Sprintf("unable to parse JSON [GET %s]: %v", c.baseURL, err))
		}
	}
	return nil
}
