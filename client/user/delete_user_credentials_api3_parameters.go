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
)

// NewDeleteUserCredentialsApi3Params creates a new DeleteUserCredentialsApi3Params object
// with the default values initialized.
func NewDeleteUserCredentialsApi3Params() *DeleteUserCredentialsApi3Params {
	var ()
	return &DeleteUserCredentialsApi3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsApi3ParamsWithTimeout creates a new DeleteUserCredentialsApi3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserCredentialsApi3ParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsApi3Params {
	var ()
	return &DeleteUserCredentialsApi3Params{

		timeout: timeout,
	}
}

// NewDeleteUserCredentialsApi3ParamsWithContext creates a new DeleteUserCredentialsApi3Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserCredentialsApi3ParamsWithContext(ctx context.Context) *DeleteUserCredentialsApi3Params {
	var ()
	return &DeleteUserCredentialsApi3Params{

		Context: ctx,
	}
}

// NewDeleteUserCredentialsApi3ParamsWithHTTPClient creates a new DeleteUserCredentialsApi3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserCredentialsApi3ParamsWithHTTPClient(client *http.Client) *DeleteUserCredentialsApi3Params {
	var ()
	return &DeleteUserCredentialsApi3Params{
		HTTPClient: client,
	}
}

/*DeleteUserCredentialsApi3Params contains all the parameters to send to the API endpoint
for the delete user credentials api3 operation typically these are written to a http.Request
*/
type DeleteUserCredentialsApi3Params struct {

	/*CredentialsApi3ID
	  id of API 3 Credential

	*/
	CredentialsApi3ID int64
	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) WithTimeout(timeout time.Duration) *DeleteUserCredentialsApi3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) WithContext(ctx context.Context) *DeleteUserCredentialsApi3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) WithHTTPClient(client *http.Client) *DeleteUserCredentialsApi3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialsApi3ID adds the credentialsApi3ID to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) WithCredentialsApi3ID(credentialsApi3ID int64) *DeleteUserCredentialsApi3Params {
	o.SetCredentialsApi3ID(credentialsApi3ID)
	return o
}

// SetCredentialsApi3ID adds the credentialsApi3Id to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) SetCredentialsApi3ID(credentialsApi3ID int64) {
	o.CredentialsApi3ID = credentialsApi3ID
}

// WithUserID adds the userID to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) WithUserID(userID int64) *DeleteUserCredentialsApi3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user credentials api3 params
func (o *DeleteUserCredentialsApi3Params) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsApi3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credentials_api3_id
	if err := r.SetPathParam("credentials_api3_id", swag.FormatInt64(o.CredentialsApi3ID)); err != nil {
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