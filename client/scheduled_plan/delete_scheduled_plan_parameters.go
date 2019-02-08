// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

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

// NewDeleteScheduledPlanParams creates a new DeleteScheduledPlanParams object
// with the default values initialized.
func NewDeleteScheduledPlanParams() *DeleteScheduledPlanParams {
	var ()
	return &DeleteScheduledPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScheduledPlanParamsWithTimeout creates a new DeleteScheduledPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScheduledPlanParamsWithTimeout(timeout time.Duration) *DeleteScheduledPlanParams {
	var ()
	return &DeleteScheduledPlanParams{

		timeout: timeout,
	}
}

// NewDeleteScheduledPlanParamsWithContext creates a new DeleteScheduledPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScheduledPlanParamsWithContext(ctx context.Context) *DeleteScheduledPlanParams {
	var ()
	return &DeleteScheduledPlanParams{

		Context: ctx,
	}
}

// NewDeleteScheduledPlanParamsWithHTTPClient creates a new DeleteScheduledPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScheduledPlanParamsWithHTTPClient(client *http.Client) *DeleteScheduledPlanParams {
	var ()
	return &DeleteScheduledPlanParams{
		HTTPClient: client,
	}
}

/*DeleteScheduledPlanParams contains all the parameters to send to the API endpoint
for the delete scheduled plan operation typically these are written to a http.Request
*/
type DeleteScheduledPlanParams struct {

	/*ScheduledPlanID
	  Scheduled Plan Id

	*/
	ScheduledPlanID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) WithTimeout(timeout time.Duration) *DeleteScheduledPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) WithContext(ctx context.Context) *DeleteScheduledPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) WithHTTPClient(client *http.Client) *DeleteScheduledPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduledPlanID adds the scheduledPlanID to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) WithScheduledPlanID(scheduledPlanID int64) *DeleteScheduledPlanParams {
	o.SetScheduledPlanID(scheduledPlanID)
	return o
}

// SetScheduledPlanID adds the scheduledPlanId to the delete scheduled plan params
func (o *DeleteScheduledPlanParams) SetScheduledPlanID(scheduledPlanID int64) {
	o.ScheduledPlanID = scheduledPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScheduledPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheduled_plan_id
	if err := r.SetPathParam("scheduled_plan_id", swag.FormatInt64(o.ScheduledPlanID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
