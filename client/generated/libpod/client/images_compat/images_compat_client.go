// Code generated by go-swagger; DO NOT EDIT.

package images_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new images compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ImageBuild(params *ImageBuildParams, opts ...ClientOption) (*ImageBuildOK, error)

	ImageCreate(params *ImageCreateParams, opts ...ClientOption) (*ImageCreateOK, error)

	ImageDelete(params *ImageDeleteParams, opts ...ClientOption) (*ImageDeleteOK, error)

	ImageGet(params *ImageGetParams, writer io.Writer, opts ...ClientOption) (*ImageGetOK, error)

	ImageGetAll(params *ImageGetAllParams, writer io.Writer, opts ...ClientOption) (*ImageGetAllOK, error)

	ImageHistory(params *ImageHistoryParams, opts ...ClientOption) (*ImageHistoryOK, error)

	ImageInspect(params *ImageInspectParams, opts ...ClientOption) (*ImageInspectOK, error)

	ImageList(params *ImageListParams, opts ...ClientOption) (*ImageListOK, error)

	ImageLoad(params *ImageLoadParams, opts ...ClientOption) (*ImageLoadOK, error)

	ImagePrune(params *ImagePruneParams, opts ...ClientOption) (*ImagePruneOK, error)

	ImagePush(params *ImagePushParams, writer io.Writer, opts ...ClientOption) (*ImagePushOK, error)

	ImageSearch(params *ImageSearchParams, opts ...ClientOption) (*ImageSearchOK, error)

	ImageTag(params *ImageTagParams, opts ...ClientOption) (*ImageTagCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ImageBuild creates image

  Build an image from the given Dockerfile(s)
*/
func (a *Client) ImageBuild(params *ImageBuildParams, opts ...ClientOption) (*ImageBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageBuild",
		Method:             "POST",
		PathPattern:        "/build",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageBuild: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageCreate creates an image

  Create an image by either pulling it from a registry or importing it.
*/
func (a *Client) ImageCreate(params *ImageCreateParams, opts ...ClientOption) (*ImageCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageCreate",
		Method:             "POST",
		PathPattern:        "/images/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageDelete removes image

  Delete an image from local storage
*/
func (a *Client) ImageDelete(params *ImageDeleteParams, opts ...ClientOption) (*ImageDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageDelete",
		Method:             "DELETE",
		PathPattern:        "/images/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageGet exports an image

  Export an image in tarball format
*/
func (a *Client) ImageGet(params *ImageGetParams, writer io.Writer, opts ...ClientOption) (*ImageGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageGet",
		Method:             "GET",
		PathPattern:        "/images/{name}/get",
		ProducesMediaTypes: []string{"application/x-tar"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageGetReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageGetAll exports several images

  Get a tarball containing all images and metadata for several image repositories
*/
func (a *Client) ImageGetAll(params *ImageGetAllParams, writer io.Writer, opts ...ClientOption) (*ImageGetAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageGetAllParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageGetAll",
		Method:             "GET",
		PathPattern:        "/images/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageGetAllReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageGetAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageGetAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageHistory histories of an image

  Return parent layers of an image.
*/
func (a *Client) ImageHistory(params *ImageHistoryParams, opts ...ClientOption) (*ImageHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageHistory",
		Method:             "GET",
		PathPattern:        "/images/{name}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageInspect inspects an image

  Return low-level information about an image.
*/
func (a *Client) ImageInspect(params *ImageInspectParams, opts ...ClientOption) (*ImageInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageInspect",
		Method:             "GET",
		PathPattern:        "/images/{name}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageList lists images

  Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.
*/
func (a *Client) ImageList(params *ImageListParams, opts ...ClientOption) (*ImageListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageList",
		Method:             "GET",
		PathPattern:        "/images/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageLoad imports image

  Load a set of images and tags into a repository.
*/
func (a *Client) ImageLoad(params *ImageLoadParams, opts ...ClientOption) (*ImageLoadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageLoadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageLoad",
		Method:             "POST",
		PathPattern:        "/images/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageLoadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageLoadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageLoad: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImagePrune prunes unused images

  Remove images from local storage that are not being used by a container
*/
func (a *Client) ImagePrune(params *ImagePruneParams, opts ...ClientOption) (*ImagePruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagePruneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImagePrune",
		Method:             "POST",
		PathPattern:        "/images/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImagePruneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImagePruneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImagePrune: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImagePush pushes image

  Push an image to a container registry
*/
func (a *Client) ImagePush(params *ImagePushParams, writer io.Writer, opts ...ClientOption) (*ImagePushOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagePushParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImagePush",
		Method:             "POST",
		PathPattern:        "/images/{name}/push",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImagePushReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImagePushOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImagePush: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageSearch searches images

  Search registries for an image
*/
func (a *Client) ImageSearch(params *ImageSearchParams, opts ...ClientOption) (*ImageSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageSearch",
		Method:             "GET",
		PathPattern:        "/images/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageTag tags an image

  Tag an image so that it becomes part of a repository.
*/
func (a *Client) ImageTag(params *ImageTagParams, opts ...ClientOption) (*ImageTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageTag",
		Method:             "POST",
		PathPattern:        "/images/{name}/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageTagReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageTagCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
