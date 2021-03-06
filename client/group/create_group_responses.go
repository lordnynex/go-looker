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

// CreateGroupReader is a Reader for the CreateGroup structure.
type CreateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGroupOK creates a CreateGroupOK with default headers values
func NewCreateGroupOK() *CreateGroupOK {
	return &CreateGroupOK{}
}

/*CreateGroupOK handles this case with default header values.

Group
*/
type CreateGroupOK struct {
	Payload *models.Group
}

func (o *CreateGroupOK) Error() string {
	return fmt.Sprintf("[POST /groups][%d] createGroupOK  %+v", 200, o.Payload)
}

func (o *CreateGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *CreateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupBadRequest creates a CreateGroupBadRequest with default headers values
func NewCreateGroupBadRequest() *CreateGroupBadRequest {
	return &CreateGroupBadRequest{}
}

/*CreateGroupBadRequest handles this case with default header values.

Bad Request
*/
type CreateGroupBadRequest struct {
	Payload *models.Error
}

func (o *CreateGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /groups][%d] createGroupBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupNotFound creates a CreateGroupNotFound with default headers values
func NewCreateGroupNotFound() *CreateGroupNotFound {
	return &CreateGroupNotFound{}
}

/*CreateGroupNotFound handles this case with default header values.

Not Found
*/
type CreateGroupNotFound struct {
	Payload *models.Error
}

func (o *CreateGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /groups][%d] createGroupNotFound  %+v", 404, o.Payload)
}

func (o *CreateGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupConflict creates a CreateGroupConflict with default headers values
func NewCreateGroupConflict() *CreateGroupConflict {
	return &CreateGroupConflict{}
}

/*CreateGroupConflict handles this case with default header values.

Resource Already Exists
*/
type CreateGroupConflict struct {
	Payload *models.Error
}

func (o *CreateGroupConflict) Error() string {
	return fmt.Sprintf("[POST /groups][%d] createGroupConflict  %+v", 409, o.Payload)
}

func (o *CreateGroupConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupUnprocessableEntity creates a CreateGroupUnprocessableEntity with default headers values
func NewCreateGroupUnprocessableEntity() *CreateGroupUnprocessableEntity {
	return &CreateGroupUnprocessableEntity{}
}

/*CreateGroupUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateGroupUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /groups][%d] createGroupUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateGroupUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
