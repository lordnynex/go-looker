// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dashboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllDashboards gets all dashboards

### Get information about all active dashboards.

Returns an array of **abbreviated dashboard objects**. Dashboards marked as deleted are excluded from this list.

Get the **full details** of a specific dashboard by id with [Dashboard](#!/Dashboard/dashboard)

Find **deleted dashboards** with [Search Dashboards](#!/Dashboard/search_dashboards)

*/
func (a *Client) AllDashboards(params *AllDashboardsParams) (*AllDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_dashboards",
		Method:             "GET",
		PathPattern:        "/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllDashboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllDashboardsOK), nil

}

/*
CreateDashboard creates dashboard

### Create a dashboard with the specified information

Creates a new dashboard object, returning the dashboard details, including the created id.

**Update** an existing dashboard with [Update Dashboard](#!/Dashboard/update_dashboard)

**Permanently delete** an existing dashboard with [Delete Dashboard](#!/Dashboard/delete_dashboard)

*/
func (a *Client) CreateDashboard(params *CreateDashboardParams) (*CreateDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_dashboard",
		Method:             "POST",
		PathPattern:        "/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDashboardOK), nil

}

/*
Dashboard gets dashboard

### Get information about the dashboard with the specified id

Returns the full details of the identified dashboard object

Get a **summary list** of all active dashboards with [All Dashboards](#!/Dashboard/all_dashboards)

**Search** for dashboards with [Search Dashboards](#!/Dashboard/search_dashboards)

*/
func (a *Client) Dashboard(params *DashboardParams) (*DashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dashboard",
		Method:             "GET",
		PathPattern:        "/dashboards/{dashboard_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DashboardOK), nil

}

/*
DeleteDashboard deletes dashboard

### Delete the dashboard with the specified id

Permanently **deletes** a dashboard. (The dashboard cannot be recovered after this operation.)

"Soft" delete or hide a dashboard by setting its `deleted` status to `True` with [Update Dashboard](#!/Dashboard/update_dashboard).

Note: When a dashboard is deleted in the UI, it is soft deleted. Use this API call to permanently remove it, if desired.

*/
func (a *Client) DeleteDashboard(params *DeleteDashboardParams) (*DeleteDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_dashboard",
		Method:             "DELETE",
		PathPattern:        "/dashboards/{dashboard_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDashboardNoContent), nil

}

/*
ImportLookmlDashboard creates dashboard

### Import a LookML dashboard to a space as a UDD
Creates a UDD (a dashboard which exists in the Looker database rather than as a LookML file) from the LookML dashboard
and puts it in the space specified. The created UDD will have a lookml_link_id which links to the original LookML dashboard.

To give the imported dashboard specify a (e.g. title: "my title") in the body of your request, otherwise the imported
dashboard will have the same title as the original LookML dashboard.

For this operation to succeed the user must have permission to see the LookML dashboard in question, and have permission to
create content in the space the dashboard is being imported to.

**Sync** a linked UDD with [Sync LookML Dashboard] (#!/Dashboard/sync_lookml_dashboard)
**Unlink** a linked UDD by setting lookml_link_id to null with [Update Dashboard](#!/Dashboard/update_dashboard)

*/
func (a *Client) ImportLookmlDashboard(params *ImportLookmlDashboardParams) (*ImportLookmlDashboardOK, *ImportLookmlDashboardCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportLookmlDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "import_lookml_dashboard",
		Method:             "POST",
		PathPattern:        "/dashboards/{lookml_dashboard_id}/import/{space_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportLookmlDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ImportLookmlDashboardOK:
		return value, nil, nil
	case *ImportLookmlDashboardCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
SearchDashboards searches dashboards

### Search all dashboards for matching criteria.

Returns an **array of dashboard objects** that match the specified search criteria.

The parameters `limit`, and `offset` are recommended for "paging" the returned results.

Get a **single dashboard** by id with [Dashboard](#!/Dashboard/dashboard)

*/
func (a *Client) SearchDashboards(params *SearchDashboardsParams) (*SearchDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "search_dashboards",
		Method:             "GET",
		PathPattern:        "/dashboards/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchDashboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchDashboardsOK), nil

}

/*
SyncLookmlDashboard updates dashboard

### Update all linked dashboards to match the specified LookML dashboard.

Any UDD (a dashboard which exists in the Looker database rather than as a LookML file) which has a lookml_link_id
which specifies the LookML dashboard's id will be updated so that it matches the current state of the LookML dashboard.

For this operation to succeed the user must have permission to view the LookML dashboard, and only linked dashboards
that the user has permission to update will be synced.

To **link** or **unlink** a UDD set the lookml_link_id with [Update Dashboard](#!/Dashboard/update_dashboard)

*/
func (a *Client) SyncLookmlDashboard(params *SyncLookmlDashboardParams) (*SyncLookmlDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncLookmlDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sync_lookml_dashboard",
		Method:             "PATCH",
		PathPattern:        "/dashboards/{lookml_dashboard_id}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SyncLookmlDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SyncLookmlDashboardOK), nil

}

/*
UpdateDashboard updates dashboard

### Update the dashboard with the specified id

Changes simple (scalar) properties of the dashboard.

*/
func (a *Client) UpdateDashboard(params *UpdateDashboardParams) (*UpdateDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_dashboard",
		Method:             "PATCH",
		PathPattern:        "/dashboards/{dashboard_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDashboardOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
