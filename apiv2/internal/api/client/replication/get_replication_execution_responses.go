// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// GetReplicationExecutionReader is a Reader for the GetReplicationExecution structure.
type GetReplicationExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationExecutionOK creates a GetReplicationExecutionOK with default headers values
func NewGetReplicationExecutionOK() *GetReplicationExecutionOK {
	return &GetReplicationExecutionOK{}
}

/*GetReplicationExecutionOK handles this case with default header values.

Success
*/
type GetReplicationExecutionOK struct {
	Payload *model.ReplicationExecution
}

func (o *GetReplicationExecutionOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionOK  %+v", 200, o.Payload)
}

func (o *GetReplicationExecutionOK) GetPayload() *model.ReplicationExecution {
	return o.Payload
}

func (o *GetReplicationExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ReplicationExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionUnauthorized creates a GetReplicationExecutionUnauthorized with default headers values
func NewGetReplicationExecutionUnauthorized() *GetReplicationExecutionUnauthorized {
	return &GetReplicationExecutionUnauthorized{}
}

/*GetReplicationExecutionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetReplicationExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationExecutionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionForbidden creates a GetReplicationExecutionForbidden with default headers values
func NewGetReplicationExecutionForbidden() *GetReplicationExecutionForbidden {
	return &GetReplicationExecutionForbidden{}
}

/*GetReplicationExecutionForbidden handles this case with default header values.

Forbidden
*/
type GetReplicationExecutionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationExecutionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionNotFound creates a GetReplicationExecutionNotFound with default headers values
func NewGetReplicationExecutionNotFound() *GetReplicationExecutionNotFound {
	return &GetReplicationExecutionNotFound{}
}

/*GetReplicationExecutionNotFound handles this case with default header values.

Not found
*/
type GetReplicationExecutionNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetReplicationExecutionNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionInternalServerError creates a GetReplicationExecutionInternalServerError with default headers values
func NewGetReplicationExecutionInternalServerError() *GetReplicationExecutionInternalServerError {
	return &GetReplicationExecutionInternalServerError{}
}

/*GetReplicationExecutionInternalServerError handles this case with default header values.

Internal server error
*/
type GetReplicationExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetReplicationExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationExecutionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
