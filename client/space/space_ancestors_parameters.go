// Code generated by go-swagger; DO NOT EDIT.

package space

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

// NewSpaceAncestorsParams creates a new SpaceAncestorsParams object
// with the default values initialized.
func NewSpaceAncestorsParams() *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSpaceAncestorsParamsWithTimeout creates a new SpaceAncestorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSpaceAncestorsParamsWithTimeout(timeout time.Duration) *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{

		timeout: timeout,
	}
}

// NewSpaceAncestorsParamsWithContext creates a new SpaceAncestorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSpaceAncestorsParamsWithContext(ctx context.Context) *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{

		Context: ctx,
	}
}

// NewSpaceAncestorsParamsWithHTTPClient creates a new SpaceAncestorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSpaceAncestorsParamsWithHTTPClient(client *http.Client) *SpaceAncestorsParams {
	var ()
	return &SpaceAncestorsParams{
		HTTPClient: client,
	}
}

/*SpaceAncestorsParams contains all the parameters to send to the API endpoint
for the space ancestors operation typically these are written to a http.Request
*/
type SpaceAncestorsParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*SpaceID
	  Id of space

	*/
	SpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the space ancestors params
func (o *SpaceAncestorsParams) WithTimeout(timeout time.Duration) *SpaceAncestorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the space ancestors params
func (o *SpaceAncestorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the space ancestors params
func (o *SpaceAncestorsParams) WithContext(ctx context.Context) *SpaceAncestorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the space ancestors params
func (o *SpaceAncestorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the space ancestors params
func (o *SpaceAncestorsParams) WithHTTPClient(client *http.Client) *SpaceAncestorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the space ancestors params
func (o *SpaceAncestorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the space ancestors params
func (o *SpaceAncestorsParams) WithFields(fields *string) *SpaceAncestorsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the space ancestors params
func (o *SpaceAncestorsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSpaceID adds the spaceID to the space ancestors params
func (o *SpaceAncestorsParams) WithSpaceID(spaceID string) *SpaceAncestorsParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the space ancestors params
func (o *SpaceAncestorsParams) SetSpaceID(spaceID string) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *SpaceAncestorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
