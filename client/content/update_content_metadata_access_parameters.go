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

// NewUpdateContentMetadataAccessParams creates a new UpdateContentMetadataAccessParams object
// with the default values initialized.
func NewUpdateContentMetadataAccessParams() *UpdateContentMetadataAccessParams {
	var ()
	return &UpdateContentMetadataAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContentMetadataAccessParamsWithTimeout creates a new UpdateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateContentMetadataAccessParamsWithTimeout(timeout time.Duration) *UpdateContentMetadataAccessParams {
	var ()
	return &UpdateContentMetadataAccessParams{

		timeout: timeout,
	}
}

// NewUpdateContentMetadataAccessParamsWithContext creates a new UpdateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateContentMetadataAccessParamsWithContext(ctx context.Context) *UpdateContentMetadataAccessParams {
	var ()
	return &UpdateContentMetadataAccessParams{

		Context: ctx,
	}
}

// NewUpdateContentMetadataAccessParamsWithHTTPClient creates a new UpdateContentMetadataAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateContentMetadataAccessParamsWithHTTPClient(client *http.Client) *UpdateContentMetadataAccessParams {
	var ()
	return &UpdateContentMetadataAccessParams{
		HTTPClient: client,
	}
}

/*UpdateContentMetadataAccessParams contains all the parameters to send to the API endpoint
for the update content metadata access operation typically these are written to a http.Request
*/
type UpdateContentMetadataAccessParams struct {

	/*Body
	  Content Metadata Access

	*/
	Body *models.ContentMetaGroupUser
	/*ContentMetadataAccessID
	  Id of content metadata access

	*/
	ContentMetadataAccessID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) WithTimeout(timeout time.Duration) *UpdateContentMetadataAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) WithContext(ctx context.Context) *UpdateContentMetadataAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) WithHTTPClient(client *http.Client) *UpdateContentMetadataAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) WithBody(body *models.ContentMetaGroupUser) *UpdateContentMetadataAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) SetBody(body *models.ContentMetaGroupUser) {
	o.Body = body
}

// WithContentMetadataAccessID adds the contentMetadataAccessID to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) WithContentMetadataAccessID(contentMetadataAccessID int64) *UpdateContentMetadataAccessParams {
	o.SetContentMetadataAccessID(contentMetadataAccessID)
	return o
}

// SetContentMetadataAccessID adds the contentMetadataAccessId to the update content metadata access params
func (o *UpdateContentMetadataAccessParams) SetContentMetadataAccessID(contentMetadataAccessID int64) {
	o.ContentMetadataAccessID = contentMetadataAccessID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContentMetadataAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param content_metadata_access_id
	if err := r.SetPathParam("content_metadata_access_id", swag.FormatInt64(o.ContentMetadataAccessID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
