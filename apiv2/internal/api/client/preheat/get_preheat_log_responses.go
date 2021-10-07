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

// GetPreheatLogReader is a Reader for the GetPreheatLog structure.
type GetPreheatLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreheatLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreheatLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPreheatLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPreheatLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPreheatLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPreheatLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPreheatLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPreheatLogOK creates a GetPreheatLogOK with default headers values
func NewGetPreheatLogOK() *GetPreheatLogOK {
	return &GetPreheatLogOK{}
}

/*GetPreheatLogOK handles this case with default header values.

Get log success
*/
type GetPreheatLogOK struct {
	/*Content type of response
	 */
	ContentType string

	Payload string
}

func (o *GetPreheatLogOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogOK  %+v", 200, o.Payload)
}

func (o *GetPreheatLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetPreheatLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreheatLogBadRequest creates a GetPreheatLogBadRequest with default headers values
func NewGetPreheatLogBadRequest() *GetPreheatLogBadRequest {
	return &GetPreheatLogBadRequest{}
}

/*GetPreheatLogBadRequest handles this case with default header values.

Bad request
*/
type GetPreheatLogBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPreheatLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetPreheatLogBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPreheatLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreheatLogUnauthorized creates a GetPreheatLogUnauthorized with default headers values
func NewGetPreheatLogUnauthorized() *GetPreheatLogUnauthorized {
	return &GetPreheatLogUnauthorized{}
}

/*GetPreheatLogUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPreheatLogUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPreheatLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPreheatLogUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPreheatLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreheatLogForbidden creates a GetPreheatLogForbidden with default headers values
func NewGetPreheatLogForbidden() *GetPreheatLogForbidden {
	return &GetPreheatLogForbidden{}
}

/*GetPreheatLogForbidden handles this case with default header values.

Forbidden
*/
type GetPreheatLogForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPreheatLogForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogForbidden  %+v", 403, o.Payload)
}

func (o *GetPreheatLogForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPreheatLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreheatLogNotFound creates a GetPreheatLogNotFound with default headers values
func NewGetPreheatLogNotFound() *GetPreheatLogNotFound {
	return &GetPreheatLogNotFound{}
}

/*GetPreheatLogNotFound handles this case with default header values.

Not found
*/
type GetPreheatLogNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPreheatLogNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogNotFound  %+v", 404, o.Payload)
}

func (o *GetPreheatLogNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPreheatLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreheatLogInternalServerError creates a GetPreheatLogInternalServerError with default headers values
func NewGetPreheatLogInternalServerError() *GetPreheatLogInternalServerError {
	return &GetPreheatLogInternalServerError{}
}

/*GetPreheatLogInternalServerError handles this case with default header values.

Internal server error
*/
type GetPreheatLogInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPreheatLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreheatLogInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPreheatLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
