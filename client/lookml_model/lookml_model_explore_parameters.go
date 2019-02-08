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

// NewLookmlModelExploreParams creates a new LookmlModelExploreParams object
// with the default values initialized.
func NewLookmlModelExploreParams() *LookmlModelExploreParams {
	var ()
	return &LookmlModelExploreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookmlModelExploreParamsWithTimeout creates a new LookmlModelExploreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookmlModelExploreParamsWithTimeout(timeout time.Duration) *LookmlModelExploreParams {
	var ()
	return &LookmlModelExploreParams{

		timeout: timeout,
	}
}

// NewLookmlModelExploreParamsWithContext creates a new LookmlModelExploreParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookmlModelExploreParamsWithContext(ctx context.Context) *LookmlModelExploreParams {
	var ()
	return &LookmlModelExploreParams{

		Context: ctx,
	}
}

// NewLookmlModelExploreParamsWithHTTPClient creates a new LookmlModelExploreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookmlModelExploreParamsWithHTTPClient(client *http.Client) *LookmlModelExploreParams {
	var ()
	return &LookmlModelExploreParams{
		HTTPClient: client,
	}
}

/*LookmlModelExploreParams contains all the parameters to send to the API endpoint
for the lookml model explore operation typically these are written to a http.Request
*/
type LookmlModelExploreParams struct {

	/*ExploreName
	  Name of explore.

	*/
	ExploreName string
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

// WithTimeout adds the timeout to the lookml model explore params
func (o *LookmlModelExploreParams) WithTimeout(timeout time.Duration) *LookmlModelExploreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookml model explore params
func (o *LookmlModelExploreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookml model explore params
func (o *LookmlModelExploreParams) WithContext(ctx context.Context) *LookmlModelExploreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookml model explore params
func (o *LookmlModelExploreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookml model explore params
func (o *LookmlModelExploreParams) WithHTTPClient(client *http.Client) *LookmlModelExploreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookml model explore params
func (o *LookmlModelExploreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExploreName adds the exploreName to the lookml model explore params
func (o *LookmlModelExploreParams) WithExploreName(exploreName string) *LookmlModelExploreParams {
	o.SetExploreName(exploreName)
	return o
}

// SetExploreName adds the exploreName to the lookml model explore params
func (o *LookmlModelExploreParams) SetExploreName(exploreName string) {
	o.ExploreName = exploreName
}

// WithFields adds the fields to the lookml model explore params
func (o *LookmlModelExploreParams) WithFields(fields *string) *LookmlModelExploreParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the lookml model explore params
func (o *LookmlModelExploreParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLookmlModelName adds the lookmlModelName to the lookml model explore params
func (o *LookmlModelExploreParams) WithLookmlModelName(lookmlModelName string) *LookmlModelExploreParams {
	o.SetLookmlModelName(lookmlModelName)
	return o
}

// SetLookmlModelName adds the lookmlModelName to the lookml model explore params
func (o *LookmlModelExploreParams) SetLookmlModelName(lookmlModelName string) {
	o.LookmlModelName = lookmlModelName
}

// WriteToRequest writes these params to a swagger request
func (o *LookmlModelExploreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param explore_name
	if err := r.SetPathParam("explore_name", o.ExploreName); err != nil {
		return err
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

	// path param lookml_model_name
	if err := r.SetPathParam("lookml_model_name", o.LookmlModelName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
