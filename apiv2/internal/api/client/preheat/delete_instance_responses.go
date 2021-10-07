// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// DeleteInstanceReader is a Reader for the DeleteInstance structure.
type DeleteInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInstanceOK creates a DeleteInstanceOK with default headers values
func NewDeleteInstanceOK() *DeleteInstanceOK {
	return &DeleteInstanceOK{}
}

/*DeleteInstanceOK handles this case with default header values.

Success
*/
type DeleteInstanceOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceOK ", 200)
}

func (o *DeleteInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteInstanceUnauthorized creates a DeleteInstanceUnauthorized with default headers values
func NewDeleteInstanceUnauthorized() *DeleteInstanceUnauthorized {
	return &DeleteInstanceUnauthorized{}
}

/*DeleteInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteInstanceUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstanceForbidden creates a DeleteInstanceForbidden with default headers values
func NewDeleteInstanceForbidden() *DeleteInstanceForbidden {
	return &DeleteInstanceForbidden{}
}

/*DeleteInstanceForbidden handles this case with default header values.

Forbidden
*/
type DeleteInstanceForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstanceForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstanceNotFound creates a DeleteInstanceNotFound with default headers values
func NewDeleteInstanceNotFound() *DeleteInstanceNotFound {
	return &DeleteInstanceNotFound{}
}

/*DeleteInstanceNotFound handles this case with default header values.

Not found
*/
type DeleteInstanceNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstanceNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstanceInternalServerError creates a DeleteInstanceInternalServerError with default headers values
func NewDeleteInstanceInternalServerError() *DeleteInstanceInternalServerError {
	return &DeleteInstanceInternalServerError{}
}

/*DeleteInstanceInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstanceInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
