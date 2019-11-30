// Code generated by go-swagger; DO NOT EDIT.

package color_collection

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

// NewColorCollectionsCustomParams creates a new ColorCollectionsCustomParams object
// with the default values initialized.
func NewColorCollectionsCustomParams() *ColorCollectionsCustomParams {
	var ()
	return &ColorCollectionsCustomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColorCollectionsCustomParamsWithTimeout creates a new ColorCollectionsCustomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColorCollectionsCustomParamsWithTimeout(timeout time.Duration) *ColorCollectionsCustomParams {
	var ()
	return &ColorCollectionsCustomParams{

		timeout: timeout,
	}
}

// NewColorCollectionsCustomParamsWithContext creates a new ColorCollectionsCustomParams object
// with the default values initialized, and the ability to set a context for a request
func NewColorCollectionsCustomParamsWithContext(ctx context.Context) *ColorCollectionsCustomParams {
	var ()
	return &ColorCollectionsCustomParams{

		Context: ctx,
	}
}

// NewColorCollectionsCustomParamsWithHTTPClient creates a new ColorCollectionsCustomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColorCollectionsCustomParamsWithHTTPClient(client *http.Client) *ColorCollectionsCustomParams {
	var ()
	return &ColorCollectionsCustomParams{
		HTTPClient: client,
	}
}

/*ColorCollectionsCustomParams contains all the parameters to send to the API endpoint
for the color collections custom operation typically these are written to a http.Request
*/
type ColorCollectionsCustomParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the color collections custom params
func (o *ColorCollectionsCustomParams) WithTimeout(timeout time.Duration) *ColorCollectionsCustomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the color collections custom params
func (o *ColorCollectionsCustomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the color collections custom params
func (o *ColorCollectionsCustomParams) WithContext(ctx context.Context) *ColorCollectionsCustomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the color collections custom params
func (o *ColorCollectionsCustomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the color collections custom params
func (o *ColorCollectionsCustomParams) WithHTTPClient(client *http.Client) *ColorCollectionsCustomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the color collections custom params
func (o *ColorCollectionsCustomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the color collections custom params
func (o *ColorCollectionsCustomParams) WithFields(fields *string) *ColorCollectionsCustomParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the color collections custom params
func (o *ColorCollectionsCustomParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ColorCollectionsCustomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
