// Code generated by go-swagger; DO NOT EDIT.

package auth

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

	models "github.com/chenrui333/go-looker/models"
)

// NewTestLdapConfigUserInfoParams creates a new TestLdapConfigUserInfoParams object
// with the default values initialized.
func NewTestLdapConfigUserInfoParams() *TestLdapConfigUserInfoParams {
	var ()
	return &TestLdapConfigUserInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigUserInfoParamsWithTimeout creates a new TestLdapConfigUserInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestLdapConfigUserInfoParamsWithTimeout(timeout time.Duration) *TestLdapConfigUserInfoParams {
	var ()
	return &TestLdapConfigUserInfoParams{

		timeout: timeout,
	}
}

// NewTestLdapConfigUserInfoParamsWithContext creates a new TestLdapConfigUserInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestLdapConfigUserInfoParamsWithContext(ctx context.Context) *TestLdapConfigUserInfoParams {
	var ()
	return &TestLdapConfigUserInfoParams{

		Context: ctx,
	}
}

// NewTestLdapConfigUserInfoParamsWithHTTPClient creates a new TestLdapConfigUserInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestLdapConfigUserInfoParamsWithHTTPClient(client *http.Client) *TestLdapConfigUserInfoParams {
	var ()
	return &TestLdapConfigUserInfoParams{
		HTTPClient: client,
	}
}

/*TestLdapConfigUserInfoParams contains all the parameters to send to the API endpoint
for the test ldap config user info operation typically these are written to a http.Request
*/
type TestLdapConfigUserInfoParams struct {

	/*Body
	  LDAP Config

	*/
	Body *models.LDAPConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) WithTimeout(timeout time.Duration) *TestLdapConfigUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) WithContext(ctx context.Context) *TestLdapConfigUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) WithHTTPClient(client *http.Client) *TestLdapConfigUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) WithBody(body *models.LDAPConfig) *TestLdapConfigUserInfoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ldap config user info params
func (o *TestLdapConfigUserInfoParams) SetBody(body *models.LDAPConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
