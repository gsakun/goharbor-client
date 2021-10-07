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

// SetCliSecretReader is a Reader for the SetCliSecret structure.
type SetCliSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCliSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetCliSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetCliSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetCliSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetCliSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetCliSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewSetCliSecretPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetCliSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetCliSecretOK creates a SetCliSecretOK with default headers values
func NewSetCliSecretOK() *SetCliSecretOK {
	return &SetCliSecretOK{}
}

/*SetCliSecretOK handles this case with default header values.

The secret is successfully updated
*/
type SetCliSecretOK struct {
}

func (o *SetCliSecretOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretOK ", 200)
}

func (o *SetCliSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCliSecretBadRequest creates a SetCliSecretBadRequest with default headers values
func NewSetCliSecretBadRequest() *SetCliSecretBadRequest {
	return &SetCliSecretBadRequest{}
}

/*SetCliSecretBadRequest handles this case with default header values.

Invalid user ID.  Or user is not onboarded via OIDC authentication. Or the secret does not meet the standard.
*/
type SetCliSecretBadRequest struct {
}

func (o *SetCliSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretBadRequest ", 400)
}

func (o *SetCliSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCliSecretUnauthorized creates a SetCliSecretUnauthorized with default headers values
func NewSetCliSecretUnauthorized() *SetCliSecretUnauthorized {
	return &SetCliSecretUnauthorized{}
}

/*SetCliSecretUnauthorized handles this case with default header values.

Unauthorized
*/
type SetCliSecretUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetCliSecretUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *SetCliSecretUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetCliSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetCliSecretForbidden creates a SetCliSecretForbidden with default headers values
func NewSetCliSecretForbidden() *SetCliSecretForbidden {
	return &SetCliSecretForbidden{}
}

/*SetCliSecretForbidden handles this case with default header values.

Forbidden
*/
type SetCliSecretForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetCliSecretForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretForbidden  %+v", 403, o.Payload)
}

func (o *SetCliSecretForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetCliSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetCliSecretNotFound creates a SetCliSecretNotFound with default headers values
func NewSetCliSecretNotFound() *SetCliSecretNotFound {
	return &SetCliSecretNotFound{}
}

/*SetCliSecretNotFound handles this case with default header values.

Not found
*/
type SetCliSecretNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetCliSecretNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretNotFound  %+v", 404, o.Payload)
}

func (o *SetCliSecretNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetCliSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetCliSecretPreconditionFailed creates a SetCliSecretPreconditionFailed with default headers values
func NewSetCliSecretPreconditionFailed() *SetCliSecretPreconditionFailed {
	return &SetCliSecretPreconditionFailed{}
}

/*SetCliSecretPreconditionFailed handles this case with default header values.

The auth mode of the system is not "oidc_auth", or the user is not onboarded via OIDC AuthN.
*/
type SetCliSecretPreconditionFailed struct {
}

func (o *SetCliSecretPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretPreconditionFailed ", 412)
}

func (o *SetCliSecretPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCliSecretInternalServerError creates a SetCliSecretInternalServerError with default headers values
func NewSetCliSecretInternalServerError() *SetCliSecretInternalServerError {
	return &SetCliSecretInternalServerError{}
}

/*SetCliSecretInternalServerError handles this case with default header values.

Internal server error
*/
type SetCliSecretInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *SetCliSecretInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/cli_secret][%d] setCliSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *SetCliSecretInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *SetCliSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
