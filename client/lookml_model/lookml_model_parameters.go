// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

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

// NewLookmlModelParams creates a new LookmlModelParams object
// with the default values initialized.
func NewLookmlModelParams() *LookmlModelParams {
	var ()
	return &LookmlModelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookmlModelParamsWithTimeout creates a new LookmlModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookmlModelParamsWithTimeout(timeout time.Duration) *LookmlModelParams {
	var ()
	return &LookmlModelParams{

		timeout: timeout,
	}
}

// NewLookmlModelParamsWithContext creates a new LookmlModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookmlModelParamsWithContext(ctx context.Context) *LookmlModelParams {
	var ()
	return &LookmlModelParams{

		Context: ctx,
	}
}

// NewLookmlModelParamsWithHTTPClient creates a new LookmlModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookmlModelParamsWithHTTPClient(client *http.Client) *LookmlModelParams {
	var ()
	return &LookmlModelParams{
		HTTPClient: client,
	}
}

/*LookmlModelParams contains all the parameters to send to the API endpoint
for the lookml model operation typically these are written to a http.Request
*/
type LookmlModelParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*LookmlModelName
	  Name of lookml model.

	*/
	LookmlModelName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lookml model params
func (o *LookmlModelParams) WithTimeout(timeout time.Duration) *LookmlModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookml model params
func (o *LookmlModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookml model params
func (o *LookmlModelParams) WithContext(ctx context.Context) *LookmlModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookml model params
func (o *LookmlModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookml model params
func (o *LookmlModelParams) WithHTTPClient(client *http.Client) *LookmlModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookml model params
func (o *LookmlModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the lookml model params
func (o *LookmlModelParams) WithFields(fields *string) *LookmlModelParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the lookml model params
func (o *LookmlModelParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLookmlModelName adds the lookmlModelName to the lookml model params
func (o *LookmlModelParams) WithLookmlModelName(lookmlModelName string) *LookmlModelParams {
	o.SetLookmlModelName(lookmlModelName)
	return o
}

// SetLookmlModelName adds the lookmlModelName to the lookml model params
func (o *LookmlModelParams) SetLookmlModelName(lookmlModelName string) {
	o.LookmlModelName = lookmlModelName
}

// WriteToRequest writes these params to a swagger request
func (o *LookmlModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param lookml_model_name
	if err := r.SetPathParam("lookml_model_name", o.LookmlModelName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
