// Code generated by go-swagger; DO NOT EDIT.

package content

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

	models "github.com/chenrui333/go-looker/models"
)

// NewUpdateContentMetadataParams creates a new UpdateContentMetadataParams object
// with the default values initialized.
func NewUpdateContentMetadataParams() *UpdateContentMetadataParams {
	var ()
	return &UpdateContentMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContentMetadataParamsWithTimeout creates a new UpdateContentMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateContentMetadataParamsWithTimeout(timeout time.Duration) *UpdateContentMetadataParams {
	var ()
	return &UpdateContentMetadataParams{

		timeout: timeout,
	}
}

// NewUpdateContentMetadataParamsWithContext creates a new UpdateContentMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateContentMetadataParamsWithContext(ctx context.Context) *UpdateContentMetadataParams {
	var ()
	return &UpdateContentMetadataParams{

		Context: ctx,
	}
}

// NewUpdateContentMetadataParamsWithHTTPClient creates a new UpdateContentMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateContentMetadataParamsWithHTTPClient(client *http.Client) *UpdateContentMetadataParams {
	var ()
	return &UpdateContentMetadataParams{
		HTTPClient: client,
	}
}

/*UpdateContentMetadataParams contains all the parameters to send to the API endpoint
for the update content metadata operation typically these are written to a http.Request
*/
type UpdateContentMetadataParams struct {

	/*Body
	  Content Metadata

	*/
	Body *models.ContentMeta
	/*ContentMetadataID
	  Id of content metadata

	*/
	ContentMetadataID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update content metadata params
func (o *UpdateContentMetadataParams) WithTimeout(timeout time.Duration) *UpdateContentMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update content metadata params
func (o *UpdateContentMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update content metadata params
func (o *UpdateContentMetadataParams) WithContext(ctx context.Context) *UpdateContentMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update content metadata params
func (o *UpdateContentMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update content metadata params
func (o *UpdateContentMetadataParams) WithHTTPClient(client *http.Client) *UpdateContentMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update content metadata params
func (o *UpdateContentMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update content metadata params
func (o *UpdateContentMetadataParams) WithBody(body *models.ContentMeta) *UpdateContentMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update content metadata params
func (o *UpdateContentMetadataParams) SetBody(body *models.ContentMeta) {
	o.Body = body
}

// WithContentMetadataID adds the contentMetadataID to the update content metadata params
func (o *UpdateContentMetadataParams) WithContentMetadataID(contentMetadataID int64) *UpdateContentMetadataParams {
	o.SetContentMetadataID(contentMetadataID)
	return o
}

// SetContentMetadataID adds the contentMetadataId to the update content metadata params
func (o *UpdateContentMetadataParams) SetContentMetadataID(contentMetadataID int64) {
	o.ContentMetadataID = contentMetadataID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContentMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param content_metadata_id
	if err := r.SetPathParam("content_metadata_id", swag.FormatInt64(o.ContentMetadataID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
