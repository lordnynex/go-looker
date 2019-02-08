// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteGitBranchParams creates a new DeleteGitBranchParams object
// with the default values initialized.
func NewDeleteGitBranchParams() *DeleteGitBranchParams {
	var ()
	return &DeleteGitBranchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGitBranchParamsWithTimeout creates a new DeleteGitBranchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGitBranchParamsWithTimeout(timeout time.Duration) *DeleteGitBranchParams {
	var ()
	return &DeleteGitBranchParams{

		timeout: timeout,
	}
}

// NewDeleteGitBranchParamsWithContext creates a new DeleteGitBranchParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGitBranchParamsWithContext(ctx context.Context) *DeleteGitBranchParams {
	var ()
	return &DeleteGitBranchParams{

		Context: ctx,
	}
}

// NewDeleteGitBranchParamsWithHTTPClient creates a new DeleteGitBranchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGitBranchParamsWithHTTPClient(client *http.Client) *DeleteGitBranchParams {
	var ()
	return &DeleteGitBranchParams{
		HTTPClient: client,
	}
}

/*DeleteGitBranchParams contains all the parameters to send to the API endpoint
for the delete git branch operation typically these are written to a http.Request
*/
type DeleteGitBranchParams struct {

	/*BranchName
	  Branch Name

	*/
	BranchName string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete git branch params
func (o *DeleteGitBranchParams) WithTimeout(timeout time.Duration) *DeleteGitBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete git branch params
func (o *DeleteGitBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete git branch params
func (o *DeleteGitBranchParams) WithContext(ctx context.Context) *DeleteGitBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete git branch params
func (o *DeleteGitBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete git branch params
func (o *DeleteGitBranchParams) WithHTTPClient(client *http.Client) *DeleteGitBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete git branch params
func (o *DeleteGitBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranchName adds the branchName to the delete git branch params
func (o *DeleteGitBranchParams) WithBranchName(branchName string) *DeleteGitBranchParams {
	o.SetBranchName(branchName)
	return o
}

// SetBranchName adds the branchName to the delete git branch params
func (o *DeleteGitBranchParams) SetBranchName(branchName string) {
	o.BranchName = branchName
}

// WithProjectID adds the projectID to the delete git branch params
func (o *DeleteGitBranchParams) WithProjectID(projectID string) *DeleteGitBranchParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete git branch params
func (o *DeleteGitBranchParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGitBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branch_name
	if err := r.SetPathParam("branch_name", o.BranchName); err != nil {
		return err
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
