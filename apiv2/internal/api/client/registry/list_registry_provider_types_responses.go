// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// ListRegistryProviderTypesReader is a Reader for the ListRegistryProviderTypes structure.
type ListRegistryProviderTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistryProviderTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistryProviderTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRegistryProviderTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRegistryProviderTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRegistryProviderTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRegistryProviderTypesOK creates a ListRegistryProviderTypesOK with default headers values
func NewListRegistryProviderTypesOK() *ListRegistryProviderTypesOK {
	return &ListRegistryProviderTypesOK{}
}

/*ListRegistryProviderTypesOK handles this case with default header values.

Success.
*/
type ListRegistryProviderTypesOK struct {
	Payload []string
}

func (o *ListRegistryProviderTypesOK) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListRegistryProviderTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderTypesUnauthorized creates a ListRegistryProviderTypesUnauthorized with default headers values
func NewListRegistryProviderTypesUnauthorized() *ListRegistryProviderTypesUnauthorized {
	return &ListRegistryProviderTypesUnauthorized{}
}

/*ListRegistryProviderTypesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRegistryProviderTypesUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderTypesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderTypesForbidden creates a ListRegistryProviderTypesForbidden with default headers values
func NewListRegistryProviderTypesForbidden() *ListRegistryProviderTypesForbidden {
	return &ListRegistryProviderTypesForbidden{}
}

/*ListRegistryProviderTypesForbidden handles this case with default header values.

Forbidden
*/
type ListRegistryProviderTypesForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderTypesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderTypesInternalServerError creates a ListRegistryProviderTypesInternalServerError with default headers values
func NewListRegistryProviderTypesInternalServerError() *ListRegistryProviderTypesInternalServerError {
	return &ListRegistryProviderTypesInternalServerError{}
}

/*ListRegistryProviderTypesInternalServerError handles this case with default header values.

Internal server error
*/
type ListRegistryProviderTypesInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListRegistryProviderTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderTypesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
