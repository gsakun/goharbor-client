// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// ManualPreheatReader is a Reader for the ManualPreheat structure.
type ManualPreheatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManualPreheatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewManualPreheatCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewManualPreheatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewManualPreheatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewManualPreheatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewManualPreheatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewManualPreheatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManualPreheatCreated creates a ManualPreheatCreated with default headers values
func NewManualPreheatCreated() *ManualPreheatCreated {
	return &ManualPreheatCreated{}
}

/*ManualPreheatCreated handles this case with default header values.

Created
*/
type ManualPreheatCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *ManualPreheatCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatCreated ", 201)
}

func (o *ManualPreheatCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewManualPreheatBadRequest creates a ManualPreheatBadRequest with default headers values
func NewManualPreheatBadRequest() *ManualPreheatBadRequest {
	return &ManualPreheatBadRequest{}
}

/*ManualPreheatBadRequest handles this case with default header values.

Bad request
*/
type ManualPreheatBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ManualPreheatBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatBadRequest  %+v", 400, o.Payload)
}

func (o *ManualPreheatBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ManualPreheatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManualPreheatUnauthorized creates a ManualPreheatUnauthorized with default headers values
func NewManualPreheatUnauthorized() *ManualPreheatUnauthorized {
	return &ManualPreheatUnauthorized{}
}

/*ManualPreheatUnauthorized handles this case with default header values.

Unauthorized
*/
type ManualPreheatUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ManualPreheatUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatUnauthorized  %+v", 401, o.Payload)
}

func (o *ManualPreheatUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ManualPreheatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManualPreheatForbidden creates a ManualPreheatForbidden with default headers values
func NewManualPreheatForbidden() *ManualPreheatForbidden {
	return &ManualPreheatForbidden{}
}

/*ManualPreheatForbidden handles this case with default header values.

Forbidden
*/
type ManualPreheatForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ManualPreheatForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatForbidden  %+v", 403, o.Payload)
}

func (o *ManualPreheatForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ManualPreheatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManualPreheatNotFound creates a ManualPreheatNotFound with default headers values
func NewManualPreheatNotFound() *ManualPreheatNotFound {
	return &ManualPreheatNotFound{}
}

/*ManualPreheatNotFound handles this case with default header values.

Not found
*/
type ManualPreheatNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ManualPreheatNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatNotFound  %+v", 404, o.Payload)
}

func (o *ManualPreheatNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ManualPreheatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManualPreheatInternalServerError creates a ManualPreheatInternalServerError with default headers values
func NewManualPreheatInternalServerError() *ManualPreheatInternalServerError {
	return &ManualPreheatInternalServerError{}
}

/*ManualPreheatInternalServerError handles this case with default header values.

Internal server error
*/
type ManualPreheatInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ManualPreheatInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] manualPreheatInternalServerError  %+v", 500, o.Payload)
}

func (o *ManualPreheatInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ManualPreheatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
