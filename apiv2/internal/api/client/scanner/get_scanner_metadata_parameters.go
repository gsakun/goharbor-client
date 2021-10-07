// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetScannerMetadataParams creates a new GetScannerMetadataParams object
// with the default values initialized.
func NewGetScannerMetadataParams() *GetScannerMetadataParams {
	var ()
	return &GetScannerMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScannerMetadataParamsWithTimeout creates a new GetScannerMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScannerMetadataParamsWithTimeout(timeout time.Duration) *GetScannerMetadataParams {
	var ()
	return &GetScannerMetadataParams{

		timeout: timeout,
	}
}

// NewGetScannerMetadataParamsWithContext creates a new GetScannerMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScannerMetadataParamsWithContext(ctx context.Context) *GetScannerMetadataParams {
	var ()
	return &GetScannerMetadataParams{

		Context: ctx,
	}
}

// NewGetScannerMetadataParamsWithHTTPClient creates a new GetScannerMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScannerMetadataParamsWithHTTPClient(client *http.Client) *GetScannerMetadataParams {
	var ()
	return &GetScannerMetadataParams{
		HTTPClient: client,
	}
}

/*GetScannerMetadataParams contains all the parameters to send to the API endpoint
for the get scanner metadata operation typically these are written to a http.Request
*/
type GetScannerMetadataParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*RegistrationID
	  The scanner registration identifier.

	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scanner metadata params
func (o *GetScannerMetadataParams) WithTimeout(timeout time.Duration) *GetScannerMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scanner metadata params
func (o *GetScannerMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scanner metadata params
func (o *GetScannerMetadataParams) WithContext(ctx context.Context) *GetScannerMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scanner metadata params
func (o *GetScannerMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scanner metadata params
func (o *GetScannerMetadataParams) WithHTTPClient(client *http.Client) *GetScannerMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scanner metadata params
func (o *GetScannerMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get scanner metadata params
func (o *GetScannerMetadataParams) WithXRequestID(xRequestID *string) *GetScannerMetadataParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get scanner metadata params
func (o *GetScannerMetadataParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistrationID adds the registrationID to the get scanner metadata params
func (o *GetScannerMetadataParams) WithRegistrationID(registrationID string) *GetScannerMetadataParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the get scanner metadata params
func (o *GetScannerMetadataParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScannerMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
