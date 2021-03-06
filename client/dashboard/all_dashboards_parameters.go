// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewAllDashboardsParams creates a new AllDashboardsParams object
// with the default values initialized.
func NewAllDashboardsParams() *AllDashboardsParams {
	var ()
	return &AllDashboardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllDashboardsParamsWithTimeout creates a new AllDashboardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllDashboardsParamsWithTimeout(timeout time.Duration) *AllDashboardsParams {
	var ()
	return &AllDashboardsParams{

		timeout: timeout,
	}
}

// NewAllDashboardsParamsWithContext creates a new AllDashboardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllDashboardsParamsWithContext(ctx context.Context) *AllDashboardsParams {
	var ()
	return &AllDashboardsParams{

		Context: ctx,
	}
}

// NewAllDashboardsParamsWithHTTPClient creates a new AllDashboardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllDashboardsParamsWithHTTPClient(client *http.Client) *AllDashboardsParams {
	var ()
	return &AllDashboardsParams{
		HTTPClient: client,
	}
}

/*AllDashboardsParams contains all the parameters to send to the API endpoint
for the all dashboards operation typically these are written to a http.Request
*/
type AllDashboardsParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all dashboards params
func (o *AllDashboardsParams) WithTimeout(timeout time.Duration) *AllDashboardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all dashboards params
func (o *AllDashboardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all dashboards params
func (o *AllDashboardsParams) WithContext(ctx context.Context) *AllDashboardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all dashboards params
func (o *AllDashboardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all dashboards params
func (o *AllDashboardsParams) WithHTTPClient(client *http.Client) *AllDashboardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all dashboards params
func (o *AllDashboardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all dashboards params
func (o *AllDashboardsParams) WithFields(fields *string) *AllDashboardsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all dashboards params
func (o *AllDashboardsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllDashboardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
