// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewAllWorkspacesParams creates a new AllWorkspacesParams object
// with the default values initialized.
func NewAllWorkspacesParams() *AllWorkspacesParams {

	return &AllWorkspacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllWorkspacesParamsWithTimeout creates a new AllWorkspacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllWorkspacesParamsWithTimeout(timeout time.Duration) *AllWorkspacesParams {

	return &AllWorkspacesParams{

		timeout: timeout,
	}
}

// NewAllWorkspacesParamsWithContext creates a new AllWorkspacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllWorkspacesParamsWithContext(ctx context.Context) *AllWorkspacesParams {

	return &AllWorkspacesParams{

		Context: ctx,
	}
}

// NewAllWorkspacesParamsWithHTTPClient creates a new AllWorkspacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllWorkspacesParamsWithHTTPClient(client *http.Client) *AllWorkspacesParams {

	return &AllWorkspacesParams{
		HTTPClient: client,
	}
}

/*AllWorkspacesParams contains all the parameters to send to the API endpoint
for the all workspaces operation typically these are written to a http.Request
*/
type AllWorkspacesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all workspaces params
func (o *AllWorkspacesParams) WithTimeout(timeout time.Duration) *AllWorkspacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all workspaces params
func (o *AllWorkspacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all workspaces params
func (o *AllWorkspacesParams) WithContext(ctx context.Context) *AllWorkspacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all workspaces params
func (o *AllWorkspacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all workspaces params
func (o *AllWorkspacesParams) WithHTTPClient(client *http.Client) *AllWorkspacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all workspaces params
func (o *AllWorkspacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AllWorkspacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
