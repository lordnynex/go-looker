// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewAllPermissionsParams creates a new AllPermissionsParams object
// with the default values initialized.
func NewAllPermissionsParams() *AllPermissionsParams {

	return &AllPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllPermissionsParamsWithTimeout creates a new AllPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllPermissionsParamsWithTimeout(timeout time.Duration) *AllPermissionsParams {

	return &AllPermissionsParams{

		timeout: timeout,
	}
}

// NewAllPermissionsParamsWithContext creates a new AllPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllPermissionsParamsWithContext(ctx context.Context) *AllPermissionsParams {

	return &AllPermissionsParams{

		Context: ctx,
	}
}

// NewAllPermissionsParamsWithHTTPClient creates a new AllPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllPermissionsParamsWithHTTPClient(client *http.Client) *AllPermissionsParams {

	return &AllPermissionsParams{
		HTTPClient: client,
	}
}

/*AllPermissionsParams contains all the parameters to send to the API endpoint
for the all permissions operation typically these are written to a http.Request
*/
type AllPermissionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all permissions params
func (o *AllPermissionsParams) WithTimeout(timeout time.Duration) *AllPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all permissions params
func (o *AllPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all permissions params
func (o *AllPermissionsParams) WithContext(ctx context.Context) *AllPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all permissions params
func (o *AllPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all permissions params
func (o *AllPermissionsParams) WithHTTPClient(client *http.Client) *AllPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all permissions params
func (o *AllPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AllPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
