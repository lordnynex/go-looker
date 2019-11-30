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

// NewUpdateGitBranchParams creates a new UpdateGitBranchParams object
// with the default values initialized.
func NewUpdateGitBranchParams() *UpdateGitBranchParams {
	var ()
	return &UpdateGitBranchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGitBranchParamsWithTimeout creates a new UpdateGitBranchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGitBranchParamsWithTimeout(timeout time.Duration) *UpdateGitBranchParams {
	var ()
	return &UpdateGitBranchParams{

		timeout: timeout,
	}
}

// NewUpdateGitBranchParamsWithContext creates a new UpdateGitBranchParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGitBranchParamsWithContext(ctx context.Context) *UpdateGitBranchParams {
	var ()
	return &UpdateGitBranchParams{

		Context: ctx,
	}
}

// NewUpdateGitBranchParamsWithHTTPClient creates a new UpdateGitBranchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGitBranchParamsWithHTTPClient(client *http.Client) *UpdateGitBranchParams {
	var ()
	return &UpdateGitBranchParams{
		HTTPClient: client,
	}
}

/*UpdateGitBranchParams contains all the parameters to send to the API endpoint
for the update git branch operation typically these are written to a http.Request
*/
type UpdateGitBranchParams struct {

	/*Body
	  Git Branch

	*/
	Body *models.GitBranch
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update git branch params
func (o *UpdateGitBranchParams) WithTimeout(timeout time.Duration) *UpdateGitBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update git branch params
func (o *UpdateGitBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update git branch params
func (o *UpdateGitBranchParams) WithContext(ctx context.Context) *UpdateGitBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update git branch params
func (o *UpdateGitBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update git branch params
func (o *UpdateGitBranchParams) WithHTTPClient(client *http.Client) *UpdateGitBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update git branch params
func (o *UpdateGitBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update git branch params
func (o *UpdateGitBranchParams) WithBody(body *models.GitBranch) *UpdateGitBranchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update git branch params
func (o *UpdateGitBranchParams) SetBody(body *models.GitBranch) {
	o.Body = body
}

// WithProjectID adds the projectID to the update git branch params
func (o *UpdateGitBranchParams) WithProjectID(projectID string) *UpdateGitBranchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update git branch params
func (o *UpdateGitBranchParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGitBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
