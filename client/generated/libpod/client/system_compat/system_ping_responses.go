// Code generated by go-swagger; DO NOT EDIT.

package system_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemPingReader is a Reader for the SystemPing structure.
type SystemPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemPingOK creates a SystemPingOK with default headers values
func NewSystemPingOK() *SystemPingOK {
	return &SystemPingOK{}
}

/* SystemPingOK describes a response with status code 200, with default header values.

Success
*/
type SystemPingOK struct {

	/* Max compatibility API Version the server supports
	 */
	APIVersion string

	/* Default version of docker image builder
	 */
	BuildKitVersion string

	/* always no-cache
	 */
	CacheControl string

	/* If the server is running with experimental mode enabled, always true
	 */
	DockerExperimental bool

	/* Max Podman API Version the server supports.
	Available if service is backed by Podman, therefore may be used to
	determine if talking to Podman engine or another engine

	*/
	LibpodAPIVersion string

	/* Default version of libpod image builder.
	Available if service is backed by Podman, therefore may be used to
	determine if talking to Podman engine or another engine

	*/
	LibpodBuildahVersion string

	/* always no-cache
	 */
	Pragma string

	Payload string
}

func (o *SystemPingOK) Error() string {
	return fmt.Sprintf("[GET /libpod/_ping][%d] systemPingOK  %+v", 200, o.Payload)
}
func (o *SystemPingOK) GetPayload() string {
	return o.Payload
}

func (o *SystemPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header API-Version
	hdrAPIVersion := response.GetHeader("API-Version")

	if hdrAPIVersion != "" {
		o.APIVersion = hdrAPIVersion
	}

	// hydrates response header BuildKit-Version
	hdrBuildKitVersion := response.GetHeader("BuildKit-Version")

	if hdrBuildKitVersion != "" {
		o.BuildKitVersion = hdrBuildKitVersion
	}

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Docker-Experimental
	hdrDockerExperimental := response.GetHeader("Docker-Experimental")

	if hdrDockerExperimental != "" {
		valdockerExperimental, err := swag.ConvertBool(hdrDockerExperimental)
		if err != nil {
			return errors.InvalidType("Docker-Experimental", "header", "bool", hdrDockerExperimental)
		}
		o.DockerExperimental = valdockerExperimental
	}

	// hydrates response header Libpod-API-Version
	hdrLibpodAPIVersion := response.GetHeader("Libpod-API-Version")

	if hdrLibpodAPIVersion != "" {
		o.LibpodAPIVersion = hdrLibpodAPIVersion
	}

	// hydrates response header Libpod-Buildah-Version
	hdrLibpodBuildahVersion := response.GetHeader("Libpod-Buildah-Version")

	if hdrLibpodBuildahVersion != "" {
		o.LibpodBuildahVersion = hdrLibpodBuildahVersion
	}

	// hydrates response header Pragma
	hdrPragma := response.GetHeader("Pragma")

	if hdrPragma != "" {
		o.Pragma = hdrPragma
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemPingInternalServerError creates a SystemPingInternalServerError with default headers values
func NewSystemPingInternalServerError() *SystemPingInternalServerError {
	return &SystemPingInternalServerError{}
}

/* SystemPingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemPingInternalServerError struct {
	Payload *SystemPingInternalServerErrorBody
}

func (o *SystemPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/_ping][%d] systemPingInternalServerError  %+v", 500, o.Payload)
}
func (o *SystemPingInternalServerError) GetPayload() *SystemPingInternalServerErrorBody {
	return o.Payload
}

func (o *SystemPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SystemPingInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SystemPingInternalServerErrorBody system ping internal server error body
swagger:model SystemPingInternalServerErrorBody
*/
type SystemPingInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this system ping internal server error body
func (o *SystemPingInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system ping internal server error body based on context it is used
func (o *SystemPingInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SystemPingInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemPingInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SystemPingInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
