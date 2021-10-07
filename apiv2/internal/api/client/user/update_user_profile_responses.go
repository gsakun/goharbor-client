// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// UpdateUserProfileReader is a Reader for the UpdateUserProfile structure.
type UpdateUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserProfileOK creates a UpdateUserProfileOK with default headers values
func NewUpdateUserProfileOK() *UpdateUserProfileOK {
	return &UpdateUserProfileOK{}
}

/*UpdateUserProfileOK handles this case with default header values.

Success
*/
type UpdateUserProfileOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateUserProfileOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileOK ", 200)
}

func (o *UpdateUserProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateUserProfileUnauthorized creates a UpdateUserProfileUnauthorized with default headers values
func NewUpdateUserProfileUnauthorized() *UpdateUserProfileUnauthorized {
	return &UpdateUserProfileUnauthorized{}
}

/*UpdateUserProfileUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserProfileUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserProfileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserProfileUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserProfileForbidden creates a UpdateUserProfileForbidden with default headers values
func NewUpdateUserProfileForbidden() *UpdateUserProfileForbidden {
	return &UpdateUserProfileForbidden{}
}

/*UpdateUserProfileForbidden handles this case with default header values.

Forbidden
*/
type UpdateUserProfileForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserProfileForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserProfileForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserProfileNotFound creates a UpdateUserProfileNotFound with default headers values
func NewUpdateUserProfileNotFound() *UpdateUserProfileNotFound {
	return &UpdateUserProfileNotFound{}
}

/*UpdateUserProfileNotFound handles this case with default header values.

Not found
*/
type UpdateUserProfileNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserProfileNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserProfileNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserProfileInternalServerError creates a UpdateUserProfileInternalServerError with default headers values
func NewUpdateUserProfileInternalServerError() *UpdateUserProfileInternalServerError {
	return &UpdateUserProfileInternalServerError{}
}

/*UpdateUserProfileInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateUserProfileInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateUserProfileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserProfileInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateUserProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
