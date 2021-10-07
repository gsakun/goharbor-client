// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// SetScannerOfProjectReader is a Reader for the SetScannerOfProject structure.
type SetScannerOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetScannerOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetScannerOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetScannerOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetScannerOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetScannerOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetScannerOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetScannerOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetScannerOfProjectOK creates a SetScannerOfProjectOK with default headers values
func NewSetScannerOfProjectOK() *SetScannerOfProjectOK {
	return &SetScannerOfProjectOK{}
}

/*SetScannerOfProjectOK handles this case with default header values.

Success
*/
type SetScannerOfProjectOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *SetScannerOfProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectOK ", 200)
}

func (o *SetScannerOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewSetScannerOfProjectBadRequest creates a SetScannerOfProjectBadRequest with default headers values
func NewSetScannerOfProjectBadRequest() *SetScannerOfProjectBadRequest {
	return &SetScannerOfProjectBadRequest{}
}

/*SetScannerOfProjectBadRequest handles this case with default header values.

Bad request
*/
type SetScannerOfProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetScannerOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *SetScannerOfProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetScannerOfProjectUnauthorized creates a SetScannerOfProjectUnauthorized with default headers values
func NewSetScannerOfProjectUnauthorized() *SetScannerOfProjectUnauthorized {
	return &SetScannerOfProjectUnauthorized{}
}

/*SetScannerOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type SetScannerOfProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetScannerOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *SetScannerOfProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetScannerOfProjectForbidden creates a SetScannerOfProjectForbidden with default headers values
func NewSetScannerOfProjectForbidden() *SetScannerOfProjectForbidden {
	return &SetScannerOfProjectForbidden{}
}

/*SetScannerOfProjectForbidden handles this case with default header values.

Forbidden
*/
type SetScannerOfProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetScannerOfProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *SetScannerOfProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetScannerOfProjectNotFound creates a SetScannerOfProjectNotFound with default headers values
func NewSetScannerOfProjectNotFound() *SetScannerOfProjectNotFound {
	return &SetScannerOfProjectNotFound{}
}

/*SetScannerOfProjectNotFound handles this case with default header values.

Not found
*/
type SetScannerOfProjectNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetScannerOfProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *SetScannerOfProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetScannerOfProjectInternalServerError creates a SetScannerOfProjectInternalServerError with default headers values
func NewSetScannerOfProjectInternalServerError() *SetScannerOfProjectInternalServerError {
	return &SetScannerOfProjectInternalServerError{}
}

/*SetScannerOfProjectInternalServerError handles this case with default header values.

Internal server error
*/
type SetScannerOfProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetScannerOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/scanner][%d] setScannerOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *SetScannerOfProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetScannerOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
