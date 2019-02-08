// Code generated by go-swagger; DO NOT EDIT.

package user

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

	models "github.com/chenrui333/go-looker/models"
)

// NewSetUserAttributeUserValueParams creates a new SetUserAttributeUserValueParams object
// with the default values initialized.
func NewSetUserAttributeUserValueParams() *SetUserAttributeUserValueParams {
	var ()
	return &SetUserAttributeUserValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetUserAttributeUserValueParamsWithTimeout creates a new SetUserAttributeUserValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetUserAttributeUserValueParamsWithTimeout(timeout time.Duration) *SetUserAttributeUserValueParams {
	var ()
	return &SetUserAttributeUserValueParams{

		timeout: timeout,
	}
}

// NewSetUserAttributeUserValueParamsWithContext creates a new SetUserAttributeUserValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetUserAttributeUserValueParamsWithContext(ctx context.Context) *SetUserAttributeUserValueParams {
	var ()
	return &SetUserAttributeUserValueParams{

		Context: ctx,
	}
}

// NewSetUserAttributeUserValueParamsWithHTTPClient creates a new SetUserAttributeUserValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetUserAttributeUserValueParamsWithHTTPClient(client *http.Client) *SetUserAttributeUserValueParams {
	var ()
	return &SetUserAttributeUserValueParams{
		HTTPClient: client,
	}
}

/*SetUserAttributeUserValueParams contains all the parameters to send to the API endpoint
for the set user attribute user value operation typically these are written to a http.Request
*/
type SetUserAttributeUserValueParams struct {

	/*Body
	  New attribute value for user.

	*/
	Body *models.UserAttributeWithValue
	/*UserAttributeID
	  Id of user attribute

	*/
	UserAttributeID int64
	/*UserID
	  Id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithTimeout(timeout time.Duration) *SetUserAttributeUserValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithContext(ctx context.Context) *SetUserAttributeUserValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithHTTPClient(client *http.Client) *SetUserAttributeUserValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithBody(body *models.UserAttributeWithValue) *SetUserAttributeUserValueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetBody(body *models.UserAttributeWithValue) {
	o.Body = body
}

// WithUserAttributeID adds the userAttributeID to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithUserAttributeID(userAttributeID int64) *SetUserAttributeUserValueParams {
	o.SetUserAttributeID(userAttributeID)
	return o
}

// SetUserAttributeID adds the userAttributeId to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetUserAttributeID(userAttributeID int64) {
	o.UserAttributeID = userAttributeID
}

// WithUserID adds the userID to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) WithUserID(userID int64) *SetUserAttributeUserValueParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set user attribute user value params
func (o *SetUserAttributeUserValueParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetUserAttributeUserValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user_attribute_id
	if err := r.SetPathParam("user_attribute_id", swag.FormatInt64(o.UserAttributeID)); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
