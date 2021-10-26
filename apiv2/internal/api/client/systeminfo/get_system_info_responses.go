// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// GetSystemInfoReader is a Reader for the GetSystemInfo structure.
type GetSystemInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSystemInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemInfoOK creates a GetSystemInfoOK with default headers values
func NewGetSystemInfoOK() *GetSystemInfoOK {
	return &GetSystemInfoOK{}
}

/*GetSystemInfoOK handles this case with default header values.

Get general info successfully.
*/
type GetSystemInfoOK struct {
	Payload *model.GeneralInfo
}

func (o *GetSystemInfoOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoOK  %+v", 200, o.Payload)
}

func (o *GetSystemInfoOK) GetPayload() *model.GeneralInfo {
	return o.Payload
}

func (o *GetSystemInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GeneralInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemInfoInternalServerError creates a GetSystemInfoInternalServerError with default headers values
func NewGetSystemInfoInternalServerError() *GetSystemInfoInternalServerError {
	return &GetSystemInfoInternalServerError{}
}

/*GetSystemInfoInternalServerError handles this case with default header values.

Internal server error
*/
type GetSystemInfoInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSystemInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo][%d] getSystemInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystemInfoInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSystemInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
