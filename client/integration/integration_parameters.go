// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIntegrationParams creates a new IntegrationParams object
// with the default values initialized.
func NewIntegrationParams() *IntegrationParams {
	var ()
	return &IntegrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationParamsWithTimeout creates a new IntegrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntegrationParamsWithTimeout(timeout time.Duration) *IntegrationParams {
	var ()
	return &IntegrationParams{

		timeout: timeout,
	}
}

// NewIntegrationParamsWithContext creates a new IntegrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewIntegrationParamsWithContext(ctx context.Context) *IntegrationParams {
	var ()
	return &IntegrationParams{

		Context: ctx,
	}
}

// NewIntegrationParamsWithHTTPClient creates a new IntegrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntegrationParamsWithHTTPClient(client *http.Client) *IntegrationParams {
	var ()
	return &IntegrationParams{
		HTTPClient: client,
	}
}

/*IntegrationParams contains all the parameters to send to the API endpoint
for the integration operation typically these are written to a http.Request
*/
type IntegrationParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*IntegrationID
	  Id of Integration

	*/
	IntegrationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the integration params
func (o *IntegrationParams) WithTimeout(timeout time.Duration) *IntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integration params
func (o *IntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integration params
func (o *IntegrationParams) WithContext(ctx context.Context) *IntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integration params
func (o *IntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integration params
func (o *IntegrationParams) WithHTTPClient(client *http.Client) *IntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integration params
func (o *IntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the integration params
func (o *IntegrationParams) WithFields(fields *string) *IntegrationParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the integration params
func (o *IntegrationParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIntegrationID adds the integrationID to the integration params
func (o *IntegrationParams) WithIntegrationID(integrationID int64) *IntegrationParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the integration params
func (o *IntegrationParams) SetIntegrationID(integrationID int64) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param integration_id
	if err := r.SetPathParam("integration_id", swag.FormatInt64(o.IntegrationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
