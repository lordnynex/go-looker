// Code generated by go-swagger; DO NOT EDIT.

package query

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

// NewQueryTaskResultsParams creates a new QueryTaskResultsParams object
// with the default values initialized.
func NewQueryTaskResultsParams() *QueryTaskResultsParams {
	var ()
	return &QueryTaskResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryTaskResultsParamsWithTimeout creates a new QueryTaskResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryTaskResultsParamsWithTimeout(timeout time.Duration) *QueryTaskResultsParams {
	var ()
	return &QueryTaskResultsParams{

		timeout: timeout,
	}
}

// NewQueryTaskResultsParamsWithContext creates a new QueryTaskResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryTaskResultsParamsWithContext(ctx context.Context) *QueryTaskResultsParams {
	var ()
	return &QueryTaskResultsParams{

		Context: ctx,
	}
}

// NewQueryTaskResultsParamsWithHTTPClient creates a new QueryTaskResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryTaskResultsParamsWithHTTPClient(client *http.Client) *QueryTaskResultsParams {
	var ()
	return &QueryTaskResultsParams{
		HTTPClient: client,
	}
}

/*QueryTaskResultsParams contains all the parameters to send to the API endpoint
for the query task results operation typically these are written to a http.Request
*/
type QueryTaskResultsParams struct {

	/*QueryTaskID
	  ID of the Query Task

	*/
	QueryTaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query task results params
func (o *QueryTaskResultsParams) WithTimeout(timeout time.Duration) *QueryTaskResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query task results params
func (o *QueryTaskResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query task results params
func (o *QueryTaskResultsParams) WithContext(ctx context.Context) *QueryTaskResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query task results params
func (o *QueryTaskResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query task results params
func (o *QueryTaskResultsParams) WithHTTPClient(client *http.Client) *QueryTaskResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query task results params
func (o *QueryTaskResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryTaskID adds the queryTaskID to the query task results params
func (o *QueryTaskResultsParams) WithQueryTaskID(queryTaskID string) *QueryTaskResultsParams {
	o.SetQueryTaskID(queryTaskID)
	return o
}

// SetQueryTaskID adds the queryTaskId to the query task results params
func (o *QueryTaskResultsParams) SetQueryTaskID(queryTaskID string) {
	o.QueryTaskID = queryTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryTaskResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param query_task_id
	if err := r.SetPathParam("query_task_id", o.QueryTaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}