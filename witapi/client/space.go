// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": space Resource Client
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
	"strconv"
)

// CreateSpacePayload is the space create action payload.
type CreateSpacePayload struct {
	Data *Space `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// CreateSpacePath computes a request path to the create action of space.
func CreateSpacePath() string {

	return fmt.Sprintf("/api/spaces")
}

// Create a space
func (c *Client) CreateSpace(ctx context.Context, path string, payload *CreateSpacePayload) (*http.Response, error) {
	req, err := c.NewCreateSpaceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateSpaceRequest create the request corresponding to the create action endpoint of the space resource.
func (c *Client) NewCreateSpaceRequest(ctx context.Context, path string, payload *CreateSpacePayload) (*http.Request, error) {
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

// DeleteSpacePath computes a request path to the delete action of space.
func DeleteSpacePath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s", param0)
}

// Delete a space with the given ID.
func (c *Client) DeleteSpace(ctx context.Context, path string, skipCluster *bool) (*http.Response, error) {
	req, err := c.NewDeleteSpaceRequest(ctx, path, skipCluster)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteSpaceRequest create the request corresponding to the delete action endpoint of the space resource.
func (c *Client) NewDeleteSpaceRequest(ctx context.Context, path string, skipCluster *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if skipCluster != nil {
		tmp141 := strconv.FormatBool(*skipCluster)
		values.Set("skipCluster", tmp141)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListSpacePath computes a request path to the list action of space.
func ListSpacePath() string {

	return fmt.Sprintf("/api/spaces")
}

// List spaces.
func (c *Client) ListSpace(ctx context.Context, path string, pageLimit *int, pageOffset *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewListSpaceRequest(ctx, path, pageLimit, pageOffset, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSpaceRequest create the request corresponding to the list action endpoint of the space resource.
func (c *Client) NewListSpaceRequest(ctx context.Context, path string, pageLimit *int, pageOffset *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if pageLimit != nil {
		tmp142 := strconv.Itoa(*pageLimit)
		values.Set("page[limit]", tmp142)
	}
	if pageOffset != nil {
		values.Set("page[offset]", *pageOffset)
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
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowSpacePath computes a request path to the show action of space.
func ShowSpacePath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s", param0)
}

// Retrieve space (as JSONAPI) for the given ID.
func (c *Client) ShowSpace(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewShowSpaceRequest(ctx, path, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowSpaceRequest create the request corresponding to the show action endpoint of the space resource.
func (c *Client) NewShowSpaceRequest(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
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

// UpdateSpacePayload is the space update action payload.
type UpdateSpacePayload struct {
	Data *Space `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// UpdateSpacePath computes a request path to the update action of space.
func UpdateSpacePath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s", param0)
}

// Update the space with the given ID.
func (c *Client) UpdateSpace(ctx context.Context, path string, payload *UpdateSpacePayload) (*http.Response, error) {
	req, err := c.NewUpdateSpaceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateSpaceRequest create the request corresponding to the update action endpoint of the space resource.
func (c *Client) NewUpdateSpaceRequest(ctx context.Context, path string, payload *UpdateSpacePayload) (*http.Request, error) {
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
	req, err := http.NewRequest("PATCH", u.String(), &body)
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