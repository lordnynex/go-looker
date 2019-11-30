// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// AddGroupGroupReader is a Reader for the AddGroupGroup structure.
type AddGroupGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGroupGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddGroupGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddGroupGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddGroupGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddGroupGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddGroupGroupOK creates a AddGroupGroupOK with default headers values
func NewAddGroupGroupOK() *AddGroupGroupOK {
	return &AddGroupGroupOK{}
}

/*AddGroupGroupOK handles this case with default header values.

Added group.
*/
type AddGroupGroupOK struct {
	Payload *models.Group
}

func (o *AddGroupGroupOK) Error() string {
	return fmt.Sprintf("[POST /groups/{group_id}/groups][%d] addGroupGroupOK  %+v", 200, o.Payload)
}

func (o *AddGroupGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *AddGroupGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGroupGroupBadRequest creates a AddGroupGroupBadRequest with default headers values
func NewAddGroupGroupBadRequest() *AddGroupGroupBadRequest {
	return &AddGroupGroupBadRequest{}
}

/*AddGroupGroupBadRequest handles this case with default header values.

Bad Request
*/
type AddGroupGroupBadRequest struct {
	Payload *models.Error
}

func (o *AddGroupGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups/{group_id}/groups][%d] addGroupGroupBadRequest  %+v", 400, o.Payload)
}

func (o *AddGroupGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddGroupGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGroupGroupForbidden creates a AddGroupGroupForbidden with default headers values
func NewAddGroupGroupForbidden() *AddGroupGroupForbidden {
	return &AddGroupGroupForbidden{}
}

/*AddGroupGroupForbidden handles this case with default header values.

Permission Denied
*/
type AddGroupGroupForbidden struct {
	Payload *models.Error
}

func (o *AddGroupGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /groups/{group_id}/groups][%d] addGroupGroupForbidden  %+v", 403, o.Payload)
}

func (o *AddGroupGroupForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddGroupGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGroupGroupNotFound creates a AddGroupGroupNotFound with default headers values
func NewAddGroupGroupNotFound() *AddGroupGroupNotFound {
	return &AddGroupGroupNotFound{}
}

/*AddGroupGroupNotFound handles this case with default header values.

Not Found
*/
type AddGroupGroupNotFound struct {
	Payload *models.Error
}

func (o *AddGroupGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /groups/{group_id}/groups][%d] addGroupGroupNotFound  %+v", 404, o.Payload)
}

func (o *AddGroupGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddGroupGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
