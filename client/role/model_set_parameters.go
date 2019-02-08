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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewModelSetParams creates a new ModelSetParams object
// with the default values initialized.
func NewModelSetParams() *ModelSetParams {
	var ()
	return &ModelSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModelSetParamsWithTimeout creates a new ModelSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModelSetParamsWithTimeout(timeout time.Duration) *ModelSetParams {
	var ()
	return &ModelSetParams{

		timeout: timeout,
	}
}

// NewModelSetParamsWithContext creates a new ModelSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewModelSetParamsWithContext(ctx context.Context) *ModelSetParams {
	var ()
	return &ModelSetParams{

		Context: ctx,
	}
}

// NewModelSetParamsWithHTTPClient creates a new ModelSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModelSetParamsWithHTTPClient(client *http.Client) *ModelSetParams {
	var ()
	return &ModelSetParams{
		HTTPClient: client,
	}
}

/*ModelSetParams contains all the parameters to send to the API endpoint
for the model set operation typically these are written to a http.Request
*/
type ModelSetParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*ModelSetID
	  Id of model set

	*/
	ModelSetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the model set params
func (o *ModelSetParams) WithTimeout(timeout time.Duration) *ModelSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the model set params
func (o *ModelSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the model set params
func (o *ModelSetParams) WithContext(ctx context.Context) *ModelSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the model set params
func (o *ModelSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the model set params
func (o *ModelSetParams) WithHTTPClient(client *http.Client) *ModelSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the model set params
func (o *ModelSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the model set params
func (o *ModelSetParams) WithFields(fields *string) *ModelSetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the model set params
func (o *ModelSetParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithModelSetID adds the modelSetID to the model set params
func (o *ModelSetParams) WithModelSetID(modelSetID int64) *ModelSetParams {
	o.SetModelSetID(modelSetID)
	return o
}

// SetModelSetID adds the modelSetId to the model set params
func (o *ModelSetParams) SetModelSetID(modelSetID int64) {
	o.ModelSetID = modelSetID
}

// WriteToRequest writes these params to a swagger request
func (o *ModelSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param model_set_id
	if err := r.SetPathParam("model_set_id", swag.FormatInt64(o.ModelSetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
