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

// NewAllProjectsParams creates a new AllProjectsParams object
// with the default values initialized.
func NewAllProjectsParams() *AllProjectsParams {
	var ()
	return &AllProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllProjectsParamsWithTimeout creates a new AllProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllProjectsParamsWithTimeout(timeout time.Duration) *AllProjectsParams {
	var ()
	return &AllProjectsParams{

		timeout: timeout,
	}
}

// NewAllProjectsParamsWithContext creates a new AllProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllProjectsParamsWithContext(ctx context.Context) *AllProjectsParams {
	var ()
	return &AllProjectsParams{

		Context: ctx,
	}
}

// NewAllProjectsParamsWithHTTPClient creates a new AllProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllProjectsParamsWithHTTPClient(client *http.Client) *AllProjectsParams {
	var ()
	return &AllProjectsParams{
		HTTPClient: client,
	}
}

/*AllProjectsParams contains all the parameters to send to the API endpoint
for the all projects operation typically these are written to a http.Request
*/
type AllProjectsParams struct {

	/*Fields
	  Requested fields

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all projects params
func (o *AllProjectsParams) WithTimeout(timeout time.Duration) *AllProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all projects params
func (o *AllProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all projects params
func (o *AllProjectsParams) WithContext(ctx context.Context) *AllProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all projects params
func (o *AllProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all projects params
func (o *AllProjectsParams) WithHTTPClient(client *http.Client) *AllProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all projects params
func (o *AllProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all projects params
func (o *AllProjectsParams) WithFields(fields *string) *AllProjectsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all projects params
func (o *AllProjectsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
