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

	"github.com/mittwald/goharbor-client/v4/apiv1/model"
)

// NewPostProjectsProjectIDMetadatasParams creates a new PostProjectsProjectIDMetadatasParams object
// with the default values initialized.
func NewPostProjectsProjectIDMetadatasParams() *PostProjectsProjectIDMetadatasParams {
	var ()
	return &PostProjectsProjectIDMetadatasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectsProjectIDMetadatasParamsWithTimeout creates a new PostProjectsProjectIDMetadatasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProjectsProjectIDMetadatasParamsWithTimeout(timeout time.Duration) *PostProjectsProjectIDMetadatasParams {
	var ()
	return &PostProjectsProjectIDMetadatasParams{

		timeout: timeout,
	}
}

// NewPostProjectsProjectIDMetadatasParamsWithContext creates a new PostProjectsProjectIDMetadatasParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProjectsProjectIDMetadatasParamsWithContext(ctx context.Context) *PostProjectsProjectIDMetadatasParams {
	var ()
	return &PostProjectsProjectIDMetadatasParams{

		Context: ctx,
	}
}

// NewPostProjectsProjectIDMetadatasParamsWithHTTPClient creates a new PostProjectsProjectIDMetadatasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProjectsProjectIDMetadatasParamsWithHTTPClient(client *http.Client) *PostProjectsProjectIDMetadatasParams {
	var ()
	return &PostProjectsProjectIDMetadatasParams{
		HTTPClient: client,
	}
}

/*PostProjectsProjectIDMetadatasParams contains all the parameters to send to the API endpoint
for the post projects project ID metadatas operation typically these are written to a http.Request
*/
type PostProjectsProjectIDMetadatasParams struct {

	/*Metadata
	  The metadata of project.

	*/
	Metadata *model.ProjectMetadata
	/*ProjectID
	  Selected project ID.

	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) WithTimeout(timeout time.Duration) *PostProjectsProjectIDMetadatasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) WithContext(ctx context.Context) *PostProjectsProjectIDMetadatasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) WithHTTPClient(client *http.Client) *PostProjectsProjectIDMetadatasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadata adds the metadata to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) WithMetadata(metadata *model.ProjectMetadata) *PostProjectsProjectIDMetadatasParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) SetMetadata(metadata *model.ProjectMetadata) {
	o.Metadata = metadata
}

// WithProjectID adds the projectID to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) WithProjectID(projectID int64) *PostProjectsProjectIDMetadatasParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post projects project ID metadatas params
func (o *PostProjectsProjectIDMetadatasParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectsProjectIDMetadatasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Metadata != nil {
		if err := r.SetBodyParam(o.Metadata); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
