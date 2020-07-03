// Code generated by go-swagger; DO NOT EDIT.

package products

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
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/model/v1_10_0"
)

// NewPutRegistriesIDParams creates a new PutRegistriesIDParams object
// with the default values initialized.
func NewPutRegistriesIDParams() *PutRegistriesIDParams {
	var ()
	return &PutRegistriesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRegistriesIDParamsWithTimeout creates a new PutRegistriesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRegistriesIDParamsWithTimeout(timeout time.Duration) *PutRegistriesIDParams {
	var ()
	return &PutRegistriesIDParams{

		timeout: timeout,
	}
}

// NewPutRegistriesIDParamsWithContext creates a new PutRegistriesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRegistriesIDParamsWithContext(ctx context.Context) *PutRegistriesIDParams {
	var ()
	return &PutRegistriesIDParams{

		Context: ctx,
	}
}

// NewPutRegistriesIDParamsWithHTTPClient creates a new PutRegistriesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRegistriesIDParamsWithHTTPClient(client *http.Client) *PutRegistriesIDParams {
	var ()
	return &PutRegistriesIDParams{
		HTTPClient: client,
	}
}

/*PutRegistriesIDParams contains all the parameters to send to the API endpoint
for the put registries ID operation typically these are written to a http.Request
*/
type PutRegistriesIDParams struct {

	/*ID
	  The registry's ID.

	*/
	ID int64
	/*RepoTarget
	  Updates registry.

	*/
	RepoTarget *v1_10_0.PutRegistry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put registries ID params
func (o *PutRegistriesIDParams) WithTimeout(timeout time.Duration) *PutRegistriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put registries ID params
func (o *PutRegistriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put registries ID params
func (o *PutRegistriesIDParams) WithContext(ctx context.Context) *PutRegistriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put registries ID params
func (o *PutRegistriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put registries ID params
func (o *PutRegistriesIDParams) WithHTTPClient(client *http.Client) *PutRegistriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put registries ID params
func (o *PutRegistriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put registries ID params
func (o *PutRegistriesIDParams) WithID(id int64) *PutRegistriesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put registries ID params
func (o *PutRegistriesIDParams) SetID(id int64) {
	o.ID = id
}

// WithRepoTarget adds the repoTarget to the put registries ID params
func (o *PutRegistriesIDParams) WithRepoTarget(repoTarget *v1_10_0.PutRegistry) *PutRegistriesIDParams {
	o.SetRepoTarget(repoTarget)
	return o
}

// SetRepoTarget adds the repoTarget to the put registries ID params
func (o *PutRegistriesIDParams) SetRepoTarget(repoTarget *v1_10_0.PutRegistry) {
	o.RepoTarget = repoTarget
}

// WriteToRequest writes these params to a swagger request
func (o *PutRegistriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.RepoTarget != nil {
		if err := r.SetBodyParam(o.RepoTarget); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
