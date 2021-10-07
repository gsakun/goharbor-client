// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// CreateWebhookPolicyOfProjectReader is a Reader for the CreateWebhookPolicyOfProject structure.
type CreateWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWebhookPolicyOfProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWebhookPolicyOfProjectCreated creates a CreateWebhookPolicyOfProjectCreated with default headers values
func NewCreateWebhookPolicyOfProjectCreated() *CreateWebhookPolicyOfProjectCreated {
	return &CreateWebhookPolicyOfProjectCreated{}
}

/*CreateWebhookPolicyOfProjectCreated handles this case with default header values.

Project webhook policy create successfully.
*/
type CreateWebhookPolicyOfProjectCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateWebhookPolicyOfProjectCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectCreated ", 201)
}

func (o *CreateWebhookPolicyOfProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateWebhookPolicyOfProjectBadRequest creates a CreateWebhookPolicyOfProjectBadRequest with default headers values
func NewCreateWebhookPolicyOfProjectBadRequest() *CreateWebhookPolicyOfProjectBadRequest {
	return &CreateWebhookPolicyOfProjectBadRequest{}
}

/*CreateWebhookPolicyOfProjectBadRequest handles this case with default header values.

Bad request
*/
type CreateWebhookPolicyOfProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookPolicyOfProjectUnauthorized creates a CreateWebhookPolicyOfProjectUnauthorized with default headers values
func NewCreateWebhookPolicyOfProjectUnauthorized() *CreateWebhookPolicyOfProjectUnauthorized {
	return &CreateWebhookPolicyOfProjectUnauthorized{}
}

/*CreateWebhookPolicyOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateWebhookPolicyOfProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookPolicyOfProjectForbidden creates a CreateWebhookPolicyOfProjectForbidden with default headers values
func NewCreateWebhookPolicyOfProjectForbidden() *CreateWebhookPolicyOfProjectForbidden {
	return &CreateWebhookPolicyOfProjectForbidden{}
}

/*CreateWebhookPolicyOfProjectForbidden handles this case with default header values.

Forbidden
*/
type CreateWebhookPolicyOfProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWebhookPolicyOfProjectInternalServerError creates a CreateWebhookPolicyOfProjectInternalServerError with default headers values
func NewCreateWebhookPolicyOfProjectInternalServerError() *CreateWebhookPolicyOfProjectInternalServerError {
	return &CreateWebhookPolicyOfProjectInternalServerError{}
}

/*CreateWebhookPolicyOfProjectInternalServerError handles this case with default header values.

Internal server error
*/
type CreateWebhookPolicyOfProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
