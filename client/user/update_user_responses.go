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

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*UpdateUserOK handles this case with default header values.

New state for specified user.
*/
type UpdateUserOK struct {
	Payload *models.User
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/*UpdateUserBadRequest handles this case with default header values.

Bad Request
*/
type UpdateUserBadRequest struct {
	Payload *models.Error
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] updateUserBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/*UpdateUserNotFound handles this case with default header values.

Not Found
*/
type UpdateUserNotFound struct {
	Payload *models.Error
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] updateUserNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserUnprocessableEntity creates a UpdateUserUnprocessableEntity with default headers values
func NewUpdateUserUnprocessableEntity() *UpdateUserUnprocessableEntity {
	return &UpdateUserUnprocessableEntity{}
}

/*UpdateUserUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateUserUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] updateUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
