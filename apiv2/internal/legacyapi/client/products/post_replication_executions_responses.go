// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostReplicationExecutionsReader is a Reader for the PostReplicationExecutions structure.
type PostReplicationExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReplicationExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostReplicationExecutionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostReplicationExecutionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostReplicationExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostReplicationExecutionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostReplicationExecutionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostReplicationExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostReplicationExecutionsCreated creates a PostReplicationExecutionsCreated with default headers values
func NewPostReplicationExecutionsCreated() *PostReplicationExecutionsCreated {
	return &PostReplicationExecutionsCreated{}
}

/* PostReplicationExecutionsCreated describes a response with status code 201, with default header values.

Success.
*/
type PostReplicationExecutionsCreated struct {
}

func (o *PostReplicationExecutionsCreated) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsCreated ", 201)
}

func (o *PostReplicationExecutionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReplicationExecutionsBadRequest creates a PostReplicationExecutionsBadRequest with default headers values
func NewPostReplicationExecutionsBadRequest() *PostReplicationExecutionsBadRequest {
	return &PostReplicationExecutionsBadRequest{}
}

/* PostReplicationExecutionsBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type PostReplicationExecutionsBadRequest struct {
}

func (o *PostReplicationExecutionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsBadRequest ", 400)
}

func (o *PostReplicationExecutionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReplicationExecutionsUnauthorized creates a PostReplicationExecutionsUnauthorized with default headers values
func NewPostReplicationExecutionsUnauthorized() *PostReplicationExecutionsUnauthorized {
	return &PostReplicationExecutionsUnauthorized{}
}

/* PostReplicationExecutionsUnauthorized describes a response with status code 401, with default header values.

User need to login first.
*/
type PostReplicationExecutionsUnauthorized struct {
}

func (o *PostReplicationExecutionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsUnauthorized ", 401)
}

func (o *PostReplicationExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReplicationExecutionsForbidden creates a PostReplicationExecutionsForbidden with default headers values
func NewPostReplicationExecutionsForbidden() *PostReplicationExecutionsForbidden {
	return &PostReplicationExecutionsForbidden{}
}

/* PostReplicationExecutionsForbidden describes a response with status code 403, with default header values.

User has no privilege for the operation.
*/
type PostReplicationExecutionsForbidden struct {
}

func (o *PostReplicationExecutionsForbidden) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsForbidden ", 403)
}

func (o *PostReplicationExecutionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReplicationExecutionsUnsupportedMediaType creates a PostReplicationExecutionsUnsupportedMediaType with default headers values
func NewPostReplicationExecutionsUnsupportedMediaType() *PostReplicationExecutionsUnsupportedMediaType {
	return &PostReplicationExecutionsUnsupportedMediaType{}
}

/* PostReplicationExecutionsUnsupportedMediaType describes a response with status code 415, with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostReplicationExecutionsUnsupportedMediaType struct {
}

func (o *PostReplicationExecutionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsUnsupportedMediaType ", 415)
}

func (o *PostReplicationExecutionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostReplicationExecutionsInternalServerError creates a PostReplicationExecutionsInternalServerError with default headers values
func NewPostReplicationExecutionsInternalServerError() *PostReplicationExecutionsInternalServerError {
	return &PostReplicationExecutionsInternalServerError{}
}

/* PostReplicationExecutionsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type PostReplicationExecutionsInternalServerError struct {
}

func (o *PostReplicationExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /replication/executions][%d] postReplicationExecutionsInternalServerError ", 500)
}

func (o *PostReplicationExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
