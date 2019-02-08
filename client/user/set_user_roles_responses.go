// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// SetUserRolesReader is a Reader for the SetUserRoles structure.
type SetUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetUserRolesOK creates a SetUserRolesOK with default headers values
func NewSetUserRolesOK() *SetUserRolesOK {
	return &SetUserRolesOK{}
}

/*SetUserRolesOK handles this case with default header values.

Roles of user.
*/
type SetUserRolesOK struct {
	Payload []*models.Role
}

func (o *SetUserRolesOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/roles][%d] setUserRolesOK  %+v", 200, o.Payload)
}

func (o *SetUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesBadRequest creates a SetUserRolesBadRequest with default headers values
func NewSetUserRolesBadRequest() *SetUserRolesBadRequest {
	return &SetUserRolesBadRequest{}
}

/*SetUserRolesBadRequest handles this case with default header values.

Bad Request
*/
type SetUserRolesBadRequest struct {
	Payload *models.Error
}

func (o *SetUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/roles][%d] setUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *SetUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetUserRolesNotFound creates a SetUserRolesNotFound with default headers values
func NewSetUserRolesNotFound() *SetUserRolesNotFound {
	return &SetUserRolesNotFound{}
}

/*SetUserRolesNotFound handles this case with default header values.

Not Found
*/
type SetUserRolesNotFound struct {
	Payload *models.Error
}

func (o *SetUserRolesNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/roles][%d] setUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *SetUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
