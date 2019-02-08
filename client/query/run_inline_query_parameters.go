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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// NewRunInlineQueryParams creates a new RunInlineQueryParams object
// with the default values initialized.
func NewRunInlineQueryParams() *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunInlineQueryParamsWithTimeout creates a new RunInlineQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunInlineQueryParamsWithTimeout(timeout time.Duration) *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{

		timeout: timeout,
	}
}

// NewRunInlineQueryParamsWithContext creates a new RunInlineQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunInlineQueryParamsWithContext(ctx context.Context) *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{

		Context: ctx,
	}
}

// NewRunInlineQueryParamsWithHTTPClient creates a new RunInlineQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunInlineQueryParamsWithHTTPClient(client *http.Client) *RunInlineQueryParams {
	var ()
	return &RunInlineQueryParams{
		HTTPClient: client,
	}
}

/*RunInlineQueryParams contains all the parameters to send to the API endpoint
for the run inline query operation typically these are written to a http.Request
*/
type RunInlineQueryParams struct {

	/*ApplyFormatting
	  Apply model-specified formatting to each result.

	*/
	ApplyFormatting *bool
	/*ApplyVis
	  Apply visualization options to results.

	*/
	ApplyVis *bool
	/*Body
	  inline query

	*/
	Body *models.Query
	/*Cache
	  Get results from cache if available.

	*/
	Cache *bool
	/*CacheOnly
	  Retrieve any results from cache even if the results have expired.

	*/
	CacheOnly *bool
	/*ForceProduction
	  Force use of production models even if the user is in development mode.

	*/
	ForceProduction *bool
	/*GenerateDrillLinks
	  Generate drill links (only applicable to 'json_detail' format.

	*/
	GenerateDrillLinks *bool
	/*ImageHeight
	  Render height for image formats.

	*/
	ImageHeight *int64
	/*ImageWidth
	  Render width for image formats.

	*/
	ImageWidth *int64
	/*Limit
	  Row limit (may override the limit in the saved query).

	*/
	Limit *int64
	/*PathPrefix
	  Prefix to use for drill links (url encoded).

	*/
	PathPrefix *string
	/*RebuildPdts
	  Rebuild PDTS used in query.

	*/
	RebuildPdts *bool
	/*ResultFormat
	  Format of result

	*/
	ResultFormat string
	/*ServerTableCalcs
	  Perform table calculations on query results

	*/
	ServerTableCalcs *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the run inline query params
func (o *RunInlineQueryParams) WithTimeout(timeout time.Duration) *RunInlineQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run inline query params
func (o *RunInlineQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run inline query params
func (o *RunInlineQueryParams) WithContext(ctx context.Context) *RunInlineQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run inline query params
func (o *RunInlineQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run inline query params
func (o *RunInlineQueryParams) WithHTTPClient(client *http.Client) *RunInlineQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run inline query params
func (o *RunInlineQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplyFormatting adds the applyFormatting to the run inline query params
func (o *RunInlineQueryParams) WithApplyFormatting(applyFormatting *bool) *RunInlineQueryParams {
	o.SetApplyFormatting(applyFormatting)
	return o
}

// SetApplyFormatting adds the applyFormatting to the run inline query params
func (o *RunInlineQueryParams) SetApplyFormatting(applyFormatting *bool) {
	o.ApplyFormatting = applyFormatting
}

// WithApplyVis adds the applyVis to the run inline query params
func (o *RunInlineQueryParams) WithApplyVis(applyVis *bool) *RunInlineQueryParams {
	o.SetApplyVis(applyVis)
	return o
}

// SetApplyVis adds the applyVis to the run inline query params
func (o *RunInlineQueryParams) SetApplyVis(applyVis *bool) {
	o.ApplyVis = applyVis
}

// WithBody adds the body to the run inline query params
func (o *RunInlineQueryParams) WithBody(body *models.Query) *RunInlineQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the run inline query params
func (o *RunInlineQueryParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithCache adds the cache to the run inline query params
func (o *RunInlineQueryParams) WithCache(cache *bool) *RunInlineQueryParams {
	o.SetCache(cache)
	return o
}

// SetCache adds the cache to the run inline query params
func (o *RunInlineQueryParams) SetCache(cache *bool) {
	o.Cache = cache
}

// WithCacheOnly adds the cacheOnly to the run inline query params
func (o *RunInlineQueryParams) WithCacheOnly(cacheOnly *bool) *RunInlineQueryParams {
	o.SetCacheOnly(cacheOnly)
	return o
}

// SetCacheOnly adds the cacheOnly to the run inline query params
func (o *RunInlineQueryParams) SetCacheOnly(cacheOnly *bool) {
	o.CacheOnly = cacheOnly
}

// WithForceProduction adds the forceProduction to the run inline query params
func (o *RunInlineQueryParams) WithForceProduction(forceProduction *bool) *RunInlineQueryParams {
	o.SetForceProduction(forceProduction)
	return o
}

// SetForceProduction adds the forceProduction to the run inline query params
func (o *RunInlineQueryParams) SetForceProduction(forceProduction *bool) {
	o.ForceProduction = forceProduction
}

// WithGenerateDrillLinks adds the generateDrillLinks to the run inline query params
func (o *RunInlineQueryParams) WithGenerateDrillLinks(generateDrillLinks *bool) *RunInlineQueryParams {
	o.SetGenerateDrillLinks(generateDrillLinks)
	return o
}

// SetGenerateDrillLinks adds the generateDrillLinks to the run inline query params
func (o *RunInlineQueryParams) SetGenerateDrillLinks(generateDrillLinks *bool) {
	o.GenerateDrillLinks = generateDrillLinks
}

// WithImageHeight adds the imageHeight to the run inline query params
func (o *RunInlineQueryParams) WithImageHeight(imageHeight *int64) *RunInlineQueryParams {
	o.SetImageHeight(imageHeight)
	return o
}

// SetImageHeight adds the imageHeight to the run inline query params
func (o *RunInlineQueryParams) SetImageHeight(imageHeight *int64) {
	o.ImageHeight = imageHeight
}

// WithImageWidth adds the imageWidth to the run inline query params
func (o *RunInlineQueryParams) WithImageWidth(imageWidth *int64) *RunInlineQueryParams {
	o.SetImageWidth(imageWidth)
	return o
}

// SetImageWidth adds the imageWidth to the run inline query params
func (o *RunInlineQueryParams) SetImageWidth(imageWidth *int64) {
	o.ImageWidth = imageWidth
}

// WithLimit adds the limit to the run inline query params
func (o *RunInlineQueryParams) WithLimit(limit *int64) *RunInlineQueryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the run inline query params
func (o *RunInlineQueryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPathPrefix adds the pathPrefix to the run inline query params
func (o *RunInlineQueryParams) WithPathPrefix(pathPrefix *string) *RunInlineQueryParams {
	o.SetPathPrefix(pathPrefix)
	return o
}

// SetPathPrefix adds the pathPrefix to the run inline query params
func (o *RunInlineQueryParams) SetPathPrefix(pathPrefix *string) {
	o.PathPrefix = pathPrefix
}

// WithRebuildPdts adds the rebuildPdts to the run inline query params
func (o *RunInlineQueryParams) WithRebuildPdts(rebuildPdts *bool) *RunInlineQueryParams {
	o.SetRebuildPdts(rebuildPdts)
	return o
}

// SetRebuildPdts adds the rebuildPdts to the run inline query params
func (o *RunInlineQueryParams) SetRebuildPdts(rebuildPdts *bool) {
	o.RebuildPdts = rebuildPdts
}

// WithResultFormat adds the resultFormat to the run inline query params
func (o *RunInlineQueryParams) WithResultFormat(resultFormat string) *RunInlineQueryParams {
	o.SetResultFormat(resultFormat)
	return o
}

// SetResultFormat adds the resultFormat to the run inline query params
func (o *RunInlineQueryParams) SetResultFormat(resultFormat string) {
	o.ResultFormat = resultFormat
}

// WithServerTableCalcs adds the serverTableCalcs to the run inline query params
func (o *RunInlineQueryParams) WithServerTableCalcs(serverTableCalcs *bool) *RunInlineQueryParams {
	o.SetServerTableCalcs(serverTableCalcs)
	return o
}

// SetServerTableCalcs adds the serverTableCalcs to the run inline query params
func (o *RunInlineQueryParams) SetServerTableCalcs(serverTableCalcs *bool) {
	o.ServerTableCalcs = serverTableCalcs
}

// WriteToRequest writes these params to a swagger request
func (o *RunInlineQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplyFormatting != nil {

		// query param apply_formatting
		var qrApplyFormatting bool
		if o.ApplyFormatting != nil {
			qrApplyFormatting = *o.ApplyFormatting
		}
		qApplyFormatting := swag.FormatBool(qrApplyFormatting)
		if qApplyFormatting != "" {
			if err := r.SetQueryParam("apply_formatting", qApplyFormatting); err != nil {
				return err
			}
		}

	}

	if o.ApplyVis != nil {

		// query param apply_vis
		var qrApplyVis bool
		if o.ApplyVis != nil {
			qrApplyVis = *o.ApplyVis
		}
		qApplyVis := swag.FormatBool(qrApplyVis)
		if qApplyVis != "" {
			if err := r.SetQueryParam("apply_vis", qApplyVis); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Cache != nil {

		// query param cache
		var qrCache bool
		if o.Cache != nil {
			qrCache = *o.Cache
		}
		qCache := swag.FormatBool(qrCache)
		if qCache != "" {
			if err := r.SetQueryParam("cache", qCache); err != nil {
				return err
			}
		}

	}

	if o.CacheOnly != nil {

		// query param cache_only
		var qrCacheOnly bool
		if o.CacheOnly != nil {
			qrCacheOnly = *o.CacheOnly
		}
		qCacheOnly := swag.FormatBool(qrCacheOnly)
		if qCacheOnly != "" {
			if err := r.SetQueryParam("cache_only", qCacheOnly); err != nil {
				return err
			}
		}

	}

	if o.ForceProduction != nil {

		// query param force_production
		var qrForceProduction bool
		if o.ForceProduction != nil {
			qrForceProduction = *o.ForceProduction
		}
		qForceProduction := swag.FormatBool(qrForceProduction)
		if qForceProduction != "" {
			if err := r.SetQueryParam("force_production", qForceProduction); err != nil {
				return err
			}
		}

	}

	if o.GenerateDrillLinks != nil {

		// query param generate_drill_links
		var qrGenerateDrillLinks bool
		if o.GenerateDrillLinks != nil {
			qrGenerateDrillLinks = *o.GenerateDrillLinks
		}
		qGenerateDrillLinks := swag.FormatBool(qrGenerateDrillLinks)
		if qGenerateDrillLinks != "" {
			if err := r.SetQueryParam("generate_drill_links", qGenerateDrillLinks); err != nil {
				return err
			}
		}

	}

	if o.ImageHeight != nil {

		// query param image_height
		var qrImageHeight int64
		if o.ImageHeight != nil {
			qrImageHeight = *o.ImageHeight
		}
		qImageHeight := swag.FormatInt64(qrImageHeight)
		if qImageHeight != "" {
			if err := r.SetQueryParam("image_height", qImageHeight); err != nil {
				return err
			}
		}

	}

	if o.ImageWidth != nil {

		// query param image_width
		var qrImageWidth int64
		if o.ImageWidth != nil {
			qrImageWidth = *o.ImageWidth
		}
		qImageWidth := swag.FormatInt64(qrImageWidth)
		if qImageWidth != "" {
			if err := r.SetQueryParam("image_width", qImageWidth); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.PathPrefix != nil {

		// query param path_prefix
		var qrPathPrefix string
		if o.PathPrefix != nil {
			qrPathPrefix = *o.PathPrefix
		}
		qPathPrefix := qrPathPrefix
		if qPathPrefix != "" {
			if err := r.SetQueryParam("path_prefix", qPathPrefix); err != nil {
				return err
			}
		}

	}

	if o.RebuildPdts != nil {

		// query param rebuild_pdts
		var qrRebuildPdts bool
		if o.RebuildPdts != nil {
			qrRebuildPdts = *o.RebuildPdts
		}
		qRebuildPdts := swag.FormatBool(qrRebuildPdts)
		if qRebuildPdts != "" {
			if err := r.SetQueryParam("rebuild_pdts", qRebuildPdts); err != nil {
				return err
			}
		}

	}

	// path param result_format
	if err := r.SetPathParam("result_format", o.ResultFormat); err != nil {
		return err
	}

	if o.ServerTableCalcs != nil {

		// query param server_table_calcs
		var qrServerTableCalcs bool
		if o.ServerTableCalcs != nil {
			qrServerTableCalcs = *o.ServerTableCalcs
		}
		qServerTableCalcs := swag.FormatBool(qrServerTableCalcs)
		if qServerTableCalcs != "" {
			if err := r.SetQueryParam("server_table_calcs", qServerTableCalcs); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
