// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": tracker Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// CreateTrackerPayload is the tracker create action payload.
type CreateTrackerPayload struct {
	Data *Tracker `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// CreateTrackerPath computes a request path to the create action of tracker.
func CreateTrackerPath() string {

	return fmt.Sprintf("/api/trackers")
}

// Add new tracker configuration.
func (c *Client) CreateTracker(ctx context.Context, path string, payload *CreateTrackerPayload) (*http.Response, error) {
	req, err := c.NewCreateTrackerRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTrackerRequest create the request corresponding to the create action endpoint of the tracker resource.
func (c *Client) NewCreateTrackerRequest(ctx context.Context, path string, payload *CreateTrackerPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// DeleteTrackerPath computes a request path to the delete action of tracker.
func DeleteTrackerPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/api/trackers/%s", param0)
}

// Delete tracker configuration.
func (c *Client) DeleteTracker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTrackerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTrackerRequest create the request corresponding to the delete action endpoint of the tracker resource.
func (c *Client) NewDeleteTrackerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListTrackerPath computes a request path to the list action of tracker.
func ListTrackerPath() string {

	return fmt.Sprintf("/api/trackers")
}

// List all tracker configurations.
func (c *Client) ListTracker(ctx context.Context, path string, filter *string, page *string) (*http.Response, error) {
	req, err := c.NewListTrackerRequest(ctx, path, filter, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTrackerRequest create the request corresponding to the list action endpoint of the tracker resource.
func (c *Client) NewListTrackerRequest(ctx context.Context, path string, filter *string, page *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if filter != nil {
		values.Set("filter", *filter)
	}
	if page != nil {
		values.Set("page", *page)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowTrackerPath computes a request path to the show action of tracker.
func ShowTrackerPath(id uuid.UUID) string {
	param0 := id.String()

	return fmt.Sprintf("/api/trackers/%s", param0)
}

// Retrieve tracker configuration for the given id.
func (c *Client) ShowTracker(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTrackerRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTrackerRequest create the request corresponding to the show action endpoint of the tracker resource.
func (c *Client) NewShowTrackerRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateTrackerPayload is the tracker update action payload.
type UpdateTrackerPayload struct {
	Data *Tracker `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// UpdateTrackerPath computes a request path to the update action of tracker.
func UpdateTrackerPath(id string) string {
	param0 := id

	return fmt.Sprintf("/api/trackers/%s", param0)
}

// Update tracker configuration.
func (c *Client) UpdateTracker(ctx context.Context, path string, payload *UpdateTrackerPayload) (*http.Response, error) {
	req, err := c.NewUpdateTrackerRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTrackerRequest create the request corresponding to the update action endpoint of the tracker resource.
func (c *Client) NewUpdateTrackerRequest(ctx context.Context, path string, payload *UpdateTrackerPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}