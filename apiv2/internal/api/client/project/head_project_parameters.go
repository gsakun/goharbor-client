// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewHeadProjectParams creates a new HeadProjectParams object
// with the default values initialized.
func NewHeadProjectParams() *HeadProjectParams {
	var ()
	return &HeadProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHeadProjectParamsWithTimeout creates a new HeadProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHeadProjectParamsWithTimeout(timeout time.Duration) *HeadProjectParams {
	var ()
	return &HeadProjectParams{

		timeout: timeout,
	}
}

// NewHeadProjectParamsWithContext creates a new HeadProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewHeadProjectParamsWithContext(ctx context.Context) *HeadProjectParams {
	var ()
	return &HeadProjectParams{

		Context: ctx,
	}
}

// NewHeadProjectParamsWithHTTPClient creates a new HeadProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHeadProjectParamsWithHTTPClient(client *http.Client) *HeadProjectParams {
	var ()
	return &HeadProjectParams{
		HTTPClient: client,
	}
}

/*HeadProjectParams contains all the parameters to send to the API endpoint
for the head project operation typically these are written to a http.Request
*/
type HeadProjectParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectName
	  Project name for checking exists.

	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the head project params
func (o *HeadProjectParams) WithTimeout(timeout time.Duration) *HeadProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head project params
func (o *HeadProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head project params
func (o *HeadProjectParams) WithContext(ctx context.Context) *HeadProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head project params
func (o *HeadProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head project params
func (o *HeadProjectParams) WithHTTPClient(client *http.Client) *HeadProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head project params
func (o *HeadProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the head project params
func (o *HeadProjectParams) WithXRequestID(xRequestID *string) *HeadProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the head project params
func (o *HeadProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the head project params
func (o *HeadProjectParams) WithProjectName(projectName string) *HeadProjectParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the head project params
func (o *HeadProjectParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *HeadProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param project_name
	qrProjectName := o.ProjectName
	qProjectName := qrProjectName
	if qProjectName != "" {
		if err := r.SetQueryParam("project_name", qProjectName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
