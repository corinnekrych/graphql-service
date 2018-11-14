// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": work_item_events Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// ListWorkItemEventsPath computes a request path to the list action of work_item_events.
func ListWorkItemEventsPath(wiID uuid.UUID) string {
	param0 := wiID.String()

	return fmt.Sprintf("/api/workitems/%s/events", param0)
}

// List events associated with the given work item
func (c *Client) ListWorkItemEvents(ctx context.Context, path string, revisionID *uuid.UUID, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewListWorkItemEventsRequest(ctx, path, revisionID, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListWorkItemEventsRequest create the request corresponding to the list action endpoint of the work_item_events resource.
func (c *Client) NewListWorkItemEventsRequest(ctx context.Context, path string, revisionID *uuid.UUID, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if revisionID != nil {
		tmp146 := revisionID.String()
		values.Set("revisionID", tmp146)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if ifModifiedSince != nil {

		header.Set("If-Modified-Since", *ifModifiedSince)
	}
	if ifNoneMatch != nil {

		header.Set("If-None-Match", *ifNoneMatch)
	}
	return req, nil
}