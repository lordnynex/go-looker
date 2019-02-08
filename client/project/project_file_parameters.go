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

// NewProjectFileParams creates a new ProjectFileParams object
// with the default values initialized.
func NewProjectFileParams() *ProjectFileParams {
	var ()
	return &ProjectFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectFileParamsWithTimeout creates a new ProjectFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectFileParamsWithTimeout(timeout time.Duration) *ProjectFileParams {
	var ()
	return &ProjectFileParams{

		timeout: timeout,
	}
}

// NewProjectFileParamsWithContext creates a new ProjectFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectFileParamsWithContext(ctx context.Context) *ProjectFileParams {
	var ()
	return &ProjectFileParams{

		Context: ctx,
	}
}

// NewProjectFileParamsWithHTTPClient creates a new ProjectFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectFileParamsWithHTTPClient(client *http.Client) *ProjectFileParams {
	var ()
	return &ProjectFileParams{
		HTTPClient: client,
	}
}

/*ProjectFileParams contains all the parameters to send to the API endpoint
for the project file operation typically these are written to a http.Request
*/
type ProjectFileParams struct {

	/*Fields
	  Requested fields

	*/
	Fields *string
	/*FileID
	  File Id

	*/
	FileID string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project file params
func (o *ProjectFileParams) WithTimeout(timeout time.Duration) *ProjectFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project file params
func (o *ProjectFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project file params
func (o *ProjectFileParams) WithContext(ctx context.Context) *ProjectFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project file params
func (o *ProjectFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project file params
func (o *ProjectFileParams) WithHTTPClient(client *http.Client) *ProjectFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project file params
func (o *ProjectFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the project file params
func (o *ProjectFileParams) WithFields(fields *string) *ProjectFileParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the project file params
func (o *ProjectFileParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFileID adds the fileID to the project file params
func (o *ProjectFileParams) WithFileID(fileID string) *ProjectFileParams {
	o.SetFileID(fileID)
	return o
}

// SetFileID adds the fileId to the project file params
func (o *ProjectFileParams) SetFileID(fileID string) {
	o.FileID = fileID
}

// WithProjectID adds the projectID to the project file params
func (o *ProjectFileParams) WithProjectID(projectID string) *ProjectFileParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project file params
func (o *ProjectFileParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param file_id
	qrFileID := o.FileID
	qFileID := qrFileID
	if qFileID != "" {
		if err := r.SetQueryParam("file_id", qFileID); err != nil {
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
