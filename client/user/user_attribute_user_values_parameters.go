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

// NewUserAttributeUserValuesParams creates a new UserAttributeUserValuesParams object
// with the default values initialized.
func NewUserAttributeUserValuesParams() *UserAttributeUserValuesParams {
	var ()
	return &UserAttributeUserValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAttributeUserValuesParamsWithTimeout creates a new UserAttributeUserValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAttributeUserValuesParamsWithTimeout(timeout time.Duration) *UserAttributeUserValuesParams {
	var ()
	return &UserAttributeUserValuesParams{

		timeout: timeout,
	}
}

// NewUserAttributeUserValuesParamsWithContext creates a new UserAttributeUserValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAttributeUserValuesParamsWithContext(ctx context.Context) *UserAttributeUserValuesParams {
	var ()
	return &UserAttributeUserValuesParams{

		Context: ctx,
	}
}

// NewUserAttributeUserValuesParamsWithHTTPClient creates a new UserAttributeUserValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAttributeUserValuesParamsWithHTTPClient(client *http.Client) *UserAttributeUserValuesParams {
	var ()
	return &UserAttributeUserValuesParams{
		HTTPClient: client,
	}
}

/*UserAttributeUserValuesParams contains all the parameters to send to the API endpoint
for the user attribute user values operation typically these are written to a http.Request
*/
type UserAttributeUserValuesParams struct {

	/*AllValues
	  If true, returns all values in the search path instead of just the first value found. Useful for debugging group precedence.

	*/
	AllValues *bool
	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*IncludeUnset
	  If true, returns an empty record for each requested attribute that has no user, group, or default value.

	*/
	IncludeUnset *bool
	/*UserAttributeIds
	  Specific user attributes to request. Omit or leave blank to request all user attributes.

	*/
	UserAttributeIds []int64
	/*UserID
	  Id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithTimeout(timeout time.Duration) *UserAttributeUserValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithContext(ctx context.Context) *UserAttributeUserValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithHTTPClient(client *http.Client) *UserAttributeUserValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllValues adds the allValues to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithAllValues(allValues *bool) *UserAttributeUserValuesParams {
	o.SetAllValues(allValues)
	return o
}

// SetAllValues adds the allValues to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetAllValues(allValues *bool) {
	o.AllValues = allValues
}

// WithFields adds the fields to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithFields(fields *string) *UserAttributeUserValuesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeUnset adds the includeUnset to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithIncludeUnset(includeUnset *bool) *UserAttributeUserValuesParams {
	o.SetIncludeUnset(includeUnset)
	return o
}

// SetIncludeUnset adds the includeUnset to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetIncludeUnset(includeUnset *bool) {
	o.IncludeUnset = includeUnset
}

// WithUserAttributeIds adds the userAttributeIds to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithUserAttributeIds(userAttributeIds []int64) *UserAttributeUserValuesParams {
	o.SetUserAttributeIds(userAttributeIds)
	return o
}

// SetUserAttributeIds adds the userAttributeIds to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetUserAttributeIds(userAttributeIds []int64) {
	o.UserAttributeIds = userAttributeIds
}

// WithUserID adds the userID to the user attribute user values params
func (o *UserAttributeUserValuesParams) WithUserID(userID int64) *UserAttributeUserValuesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user attribute user values params
func (o *UserAttributeUserValuesParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserAttributeUserValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllValues != nil {

		// query param all_values
		var qrAllValues bool
		if o.AllValues != nil {
			qrAllValues = *o.AllValues
		}
		qAllValues := swag.FormatBool(qrAllValues)
		if qAllValues != "" {
			if err := r.SetQueryParam("all_values", qAllValues); err != nil {
				return err
			}
		}

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

	if o.IncludeUnset != nil {

		// query param include_unset
		var qrIncludeUnset bool
		if o.IncludeUnset != nil {
			qrIncludeUnset = *o.IncludeUnset
		}
		qIncludeUnset := swag.FormatBool(qrIncludeUnset)
		if qIncludeUnset != "" {
			if err := r.SetQueryParam("include_unset", qIncludeUnset); err != nil {
				return err
			}
		}

	}

	var valuesUserAttributeIds []string
	for _, v := range o.UserAttributeIds {
		valuesUserAttributeIds = append(valuesUserAttributeIds, swag.FormatInt64(v))
	}

	joinedUserAttributeIds := swag.JoinByFormat(valuesUserAttributeIds, "csv")
	// query array param user_attribute_ids
	if err := r.SetQueryParam("user_attribute_ids", joinedUserAttributeIds...); err != nil {
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
