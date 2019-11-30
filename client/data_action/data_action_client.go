// Code generated by go-swagger; DO NOT EDIT.

package data_action

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data action API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data action API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FetchRemoteDataActionForm fetches remote data action form

For some data actions, the remote server may supply a form requesting further user input. This endpoint takes a data action, asks the remote server to generate a form for it, and returns that form to you for presentation to the user.
*/
func (a *Client) FetchRemoteDataActionForm(params *FetchRemoteDataActionFormParams) (*FetchRemoteDataActionFormOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchRemoteDataActionFormParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fetch_remote_data_action_form",
		Method:             "POST",
		PathPattern:        "/data_actions/form",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchRemoteDataActionFormReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FetchRemoteDataActionFormOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetch_remote_data_action_form: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PerformDataAction sends a data action

Perform a data action. The data action object can be obtained from query results, and used to perform an arbitrary action.
*/
func (a *Client) PerformDataAction(params *PerformDataActionParams) (*PerformDataActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformDataActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "perform_data_action",
		Method:             "POST",
		PathPattern:        "/data_actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformDataActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PerformDataActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for perform_data_action: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
