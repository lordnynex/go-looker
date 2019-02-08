// Code generated by go-swagger; DO NOT EDIT.

package query

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

	models "github.com/chenrui333/go-looker/models"
)

// NewCreateQueryParams creates a new CreateQueryParams object
// with the default values initialized.
func NewCreateQueryParams() *CreateQueryParams {
	var ()
	return &CreateQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQueryParamsWithTimeout creates a new CreateQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateQueryParamsWithTimeout(timeout time.Duration) *CreateQueryParams {
	var ()
	return &CreateQueryParams{

		timeout: timeout,
	}
}

// NewCreateQueryParamsWithContext creates a new CreateQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateQueryParamsWithContext(ctx context.Context) *CreateQueryParams {
	var ()
	return &CreateQueryParams{

		Context: ctx,
	}
}

// NewCreateQueryParamsWithHTTPClient creates a new CreateQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateQueryParamsWithHTTPClient(client *http.Client) *CreateQueryParams {
	var ()
	return &CreateQueryParams{
		HTTPClient: client,
	}
}

/*CreateQueryParams contains all the parameters to send to the API endpoint
for the create query operation typically these are written to a http.Request
*/
type CreateQueryParams struct {

	/*Body
	  Query

	*/
	Body *models.Query
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create query params
func (o *CreateQueryParams) WithTimeout(timeout time.Duration) *CreateQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create query params
func (o *CreateQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create query params
func (o *CreateQueryParams) WithContext(ctx context.Context) *CreateQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create query params
func (o *CreateQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create query params
func (o *CreateQueryParams) WithHTTPClient(client *http.Client) *CreateQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create query params
func (o *CreateQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create query params
func (o *CreateQueryParams) WithBody(body *models.Query) *CreateQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create query params
func (o *CreateQueryParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithFields adds the fields to the create query params
func (o *CreateQueryParams) WithFields(fields *string) *CreateQueryParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create query params
func (o *CreateQueryParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
