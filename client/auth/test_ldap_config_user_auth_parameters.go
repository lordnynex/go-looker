// Code generated by go-swagger; DO NOT EDIT.

package auth

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

	models "github.com/chenrui333/go-looker/models"
)

// NewTestLdapConfigUserAuthParams creates a new TestLdapConfigUserAuthParams object
// with the default values initialized.
func NewTestLdapConfigUserAuthParams() *TestLdapConfigUserAuthParams {
	var ()
	return &TestLdapConfigUserAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigUserAuthParamsWithTimeout creates a new TestLdapConfigUserAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestLdapConfigUserAuthParamsWithTimeout(timeout time.Duration) *TestLdapConfigUserAuthParams {
	var ()
	return &TestLdapConfigUserAuthParams{

		timeout: timeout,
	}
}

// NewTestLdapConfigUserAuthParamsWithContext creates a new TestLdapConfigUserAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestLdapConfigUserAuthParamsWithContext(ctx context.Context) *TestLdapConfigUserAuthParams {
	var ()
	return &TestLdapConfigUserAuthParams{

		Context: ctx,
	}
}

// NewTestLdapConfigUserAuthParamsWithHTTPClient creates a new TestLdapConfigUserAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestLdapConfigUserAuthParamsWithHTTPClient(client *http.Client) *TestLdapConfigUserAuthParams {
	var ()
	return &TestLdapConfigUserAuthParams{
		HTTPClient: client,
	}
}

/*TestLdapConfigUserAuthParams contains all the parameters to send to the API endpoint
for the test ldap config user auth operation typically these are written to a http.Request
*/
type TestLdapConfigUserAuthParams struct {

	/*Body
	  LDAP Config

	*/
	Body *models.LDAPConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) WithTimeout(timeout time.Duration) *TestLdapConfigUserAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) WithContext(ctx context.Context) *TestLdapConfigUserAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) WithHTTPClient(client *http.Client) *TestLdapConfigUserAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) WithBody(body *models.LDAPConfig) *TestLdapConfigUserAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ldap config user auth params
func (o *TestLdapConfigUserAuthParams) SetBody(body *models.LDAPConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigUserAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
