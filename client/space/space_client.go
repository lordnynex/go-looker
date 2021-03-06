// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new space API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for space API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllSpaces gets all spaces

### Get information about all spaces.
*/
func (a *Client) AllSpaces(params *AllSpacesParams) (*AllSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_spaces",
		Method:             "GET",
		PathPattern:        "/spaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllSpacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateSpace creates space

### Create a space with specified information.

Caller must have permission to edit the parent space and to create spaces, otherwise the request
returns 404 Not Found.

*/
func (a *Client) CreateSpace(params *CreateSpaceParams) (*CreateSpaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_space",
		Method:             "POST",
		PathPattern:        "/spaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSpaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSpaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_space: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteSpace deletes space

### Delete the space with a specific id including any children spaces.
**DANGER** this will delete all looks and dashboards in the space.

*/
func (a *Client) DeleteSpace(params *DeleteSpaceParams) (*DeleteSpaceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_space",
		Method:             "DELETE",
		PathPattern:        "/spaces/{space_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSpaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSpaceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_space: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchSpaces searches spaces

### Search Spaces

  Returns an **array of space objects** that match the given search criteria.

  If multiple search params are given and `filter_or` is FALSE or not specified,
search params are combined in a logical AND operation.
Only rows that match *all* search param criteria will be returned.

If `filter_or` is TRUE, multiple search params are combined in a logical OR operation.
Results will include rows that match **any** of the search criteria.

String search params use case-insensitive matching.
String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions.
example="dan%" will match "danger" and "Danzig" but not "David"
example="D_m%" will match "Damage" and "dump"

Integer search params can accept a single value or a comma separated list of values. The multiple
values will be combined under a logical OR operation - results will match at least one of
the given values.

Most search params can accept "IS NULL" and "NOT NULL" as special expressions to match
or exclude (respectively) rows where the column is null.

Boolean search params accept only "true" and "false" as values.


  The parameters `limit`, and `offset` are recommended for fetching results in page-size chunks.

  Get a **single space** by id with [Space](#!/Space/space)

*/
func (a *Client) SearchSpaces(params *SearchSpacesParams) (*SearchSpacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchSpacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "search_spaces",
		Method:             "GET",
		PathPattern:        "/spaces/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchSpacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchSpacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_spaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Space gets space

### Get information about the space with a specific id.
*/
func (a *Client) Space(params *SpaceParams) (*SpaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceAncestors gets space ancestors

### Get the ancestors of a space
*/
func (a *Client) SpaceAncestors(params *SpaceAncestorsParams) (*SpaceAncestorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceAncestorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_ancestors",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/ancestors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceAncestorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceAncestorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_ancestors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceChildren gets space children

### Get the children of a space.
*/
func (a *Client) SpaceChildren(params *SpaceChildrenParams) (*SpaceChildrenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceChildrenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_children",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceChildrenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceChildrenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_children: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceChildrenSearch searches space children

### Search the children of a space
*/
func (a *Client) SpaceChildrenSearch(params *SpaceChildrenSearchParams) (*SpaceChildrenSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceChildrenSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_children_search",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/children/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceChildrenSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceChildrenSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_children_search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceDashboards gets space dashboards

### Get the dashboards in a space
*/
func (a *Client) SpaceDashboards(params *SpaceDashboardsParams) (*SpaceDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_dashboards",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceDashboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_dashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceLooks gets space looks

### Get the looks in a space
*/
func (a *Client) SpaceLooks(params *SpaceLooksParams) (*SpaceLooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceLooksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_looks",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/looks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceLooksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceLooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_looks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SpaceParent gets space parent

### Get the parent of a space
*/
func (a *Client) SpaceParent(params *SpaceParentParams) (*SpaceParentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSpaceParentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "space_parent",
		Method:             "GET",
		PathPattern:        "/spaces/{space_id}/parent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SpaceParentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SpaceParentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for space_parent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateSpace updates space

### Update the space with a specific id.
*/
func (a *Client) UpdateSpace(params *UpdateSpaceParams) (*UpdateSpaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSpaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_space",
		Method:             "PATCH",
		PathPattern:        "/spaces/{space_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSpaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSpaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_space: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
