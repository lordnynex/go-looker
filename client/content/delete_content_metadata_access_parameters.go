// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewDeleteContentMetadataAccessParams creates a new DeleteContentMetadataAccessParams object
// with the default values initialized.
func NewDeleteContentMetadataAccessParams() *DeleteContentMetadataAccessParams {
	var ()
	return &DeleteContentMetadataAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentMetadataAccessParamsWithTimeout creates a new DeleteContentMetadataAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteContentMetadataAccessParamsWithTimeout(timeout time.Duration) *DeleteContentMetadataAccessParams {
	var ()
	return &DeleteContentMetadataAccessParams{

		timeout: timeout,
	}
}

// NewDeleteContentMetadataAccessParamsWithContext creates a new DeleteContentMetadataAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteContentMetadataAccessParamsWithContext(ctx context.Context) *DeleteContentMetadataAccessParams {
	var ()
	return &DeleteContentMetadataAccessParams{

		Context: ctx,
	}
}

// NewDeleteContentMetadataAccessParamsWithHTTPClient creates a new DeleteContentMetadataAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteContentMetadataAccessParamsWithHTTPClient(client *http.Client) *DeleteContentMetadataAccessParams {
	var ()
	return &DeleteContentMetadataAccessParams{
		HTTPClient: client,
	}
}

/*DeleteContentMetadataAccessParams contains all the parameters to send to the API endpoint
for the delete content metadata access operation typically these are written to a http.Request
*/
type DeleteContentMetadataAccessParams struct {

	/*ContentMetadataAccessID
	  Id of content metadata access

	*/
	ContentMetadataAccessID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) WithTimeout(timeout time.Duration) *DeleteContentMetadataAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) WithContext(ctx context.Context) *DeleteContentMetadataAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) WithHTTPClient(client *http.Client) *DeleteContentMetadataAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentMetadataAccessID adds the contentMetadataAccessID to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) WithContentMetadataAccessID(contentMetadataAccessID int64) *DeleteContentMetadataAccessParams {
	o.SetContentMetadataAccessID(contentMetadataAccessID)
	return o
}

// SetContentMetadataAccessID adds the contentMetadataAccessId to the delete content metadata access params
func (o *DeleteContentMetadataAccessParams) SetContentMetadataAccessID(contentMetadataAccessID int64) {
	o.ContentMetadataAccessID = contentMetadataAccessID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentMetadataAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param content_metadata_access_id
	if err := r.SetPathParam("content_metadata_access_id", swag.FormatInt64(o.ContentMetadataAccessID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}