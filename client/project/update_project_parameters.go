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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// NewUpdateProjectParams creates a new UpdateProjectParams object
// with the default values initialized.
func NewUpdateProjectParams() *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectParamsWithTimeout creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProjectParamsWithTimeout(timeout time.Duration) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		timeout: timeout,
	}
}

// NewUpdateProjectParamsWithContext creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProjectParamsWithContext(ctx context.Context) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		Context: ctx,
	}
}

// NewUpdateProjectParamsWithHTTPClient creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProjectParamsWithHTTPClient(client *http.Client) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{
		HTTPClient: client,
	}
}

/*UpdateProjectParams contains all the parameters to send to the API endpoint
for the update project operation typically these are written to a http.Request
*/
type UpdateProjectParams struct {

	/*Body
	  Project

	*/
	Body *models.Project
	/*Fields
	  Requested fields

	*/
	Fields *string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) WithTimeout(timeout time.Duration) *UpdateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project params
func (o *UpdateProjectParams) WithContext(ctx context.Context) *UpdateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project params
func (o *UpdateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) WithHTTPClient(client *http.Client) *UpdateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update project params
func (o *UpdateProjectParams) WithBody(body *models.Project) *UpdateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project params
func (o *UpdateProjectParams) SetBody(body *models.Project) {
	o.Body = body
}

// WithFields adds the fields to the update project params
func (o *UpdateProjectParams) WithFields(fields *string) *UpdateProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update project params
func (o *UpdateProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectID adds the projectID to the update project params
func (o *UpdateProjectParams) WithProjectID(projectID string) *UpdateProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update project params
func (o *UpdateProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
