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
)

// NewAllLookmlTestsParams creates a new AllLookmlTestsParams object
// with the default values initialized.
func NewAllLookmlTestsParams() *AllLookmlTestsParams {
	var ()
	return &AllLookmlTestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllLookmlTestsParamsWithTimeout creates a new AllLookmlTestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllLookmlTestsParamsWithTimeout(timeout time.Duration) *AllLookmlTestsParams {
	var ()
	return &AllLookmlTestsParams{

		timeout: timeout,
	}
}

// NewAllLookmlTestsParamsWithContext creates a new AllLookmlTestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllLookmlTestsParamsWithContext(ctx context.Context) *AllLookmlTestsParams {
	var ()
	return &AllLookmlTestsParams{

		Context: ctx,
	}
}

// NewAllLookmlTestsParamsWithHTTPClient creates a new AllLookmlTestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllLookmlTestsParamsWithHTTPClient(client *http.Client) *AllLookmlTestsParams {
	var ()
	return &AllLookmlTestsParams{
		HTTPClient: client,
	}
}

/*AllLookmlTestsParams contains all the parameters to send to the API endpoint
for the all lookml tests operation typically these are written to a http.Request
*/
type AllLookmlTestsParams struct {

	/*FileID
	  File Id

	*/
	FileID *string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all lookml tests params
func (o *AllLookmlTestsParams) WithTimeout(timeout time.Duration) *AllLookmlTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all lookml tests params
func (o *AllLookmlTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all lookml tests params
func (o *AllLookmlTestsParams) WithContext(ctx context.Context) *AllLookmlTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all lookml tests params
func (o *AllLookmlTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all lookml tests params
func (o *AllLookmlTestsParams) WithHTTPClient(client *http.Client) *AllLookmlTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all lookml tests params
func (o *AllLookmlTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileID adds the fileID to the all lookml tests params
func (o *AllLookmlTestsParams) WithFileID(fileID *string) *AllLookmlTestsParams {
	o.SetFileID(fileID)
	return o
}

// SetFileID adds the fileId to the all lookml tests params
func (o *AllLookmlTestsParams) SetFileID(fileID *string) {
	o.FileID = fileID
}

// WithProjectID adds the projectID to the all lookml tests params
func (o *AllLookmlTestsParams) WithProjectID(projectID string) *AllLookmlTestsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the all lookml tests params
func (o *AllLookmlTestsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *AllLookmlTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileID != nil {

		// query param file_id
		var qrFileID string
		if o.FileID != nil {
			qrFileID = *o.FileID
		}
		qFileID := qrFileID
		if qFileID != "" {
			if err := r.SetQueryParam("file_id", qFileID); err != nil {
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
