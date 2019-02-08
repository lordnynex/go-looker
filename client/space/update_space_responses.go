// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdateSpaceReader is a Reader for the UpdateSpace structure.
type UpdateSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSpaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateSpaceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSpaceOK creates a UpdateSpaceOK with default headers values
func NewUpdateSpaceOK() *UpdateSpaceOK {
	return &UpdateSpaceOK{}
}

/*UpdateSpaceOK handles this case with default header values.

Space
*/
type UpdateSpaceOK struct {
	Payload *models.Space
}

func (o *UpdateSpaceOK) Error() string {
	return fmt.Sprintf("[PATCH /spaces/{space_id}][%d] updateSpaceOK  %+v", 200, o.Payload)
}

func (o *UpdateSpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Space)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpaceBadRequest creates a UpdateSpaceBadRequest with default headers values
func NewUpdateSpaceBadRequest() *UpdateSpaceBadRequest {
	return &UpdateSpaceBadRequest{}
}

/*UpdateSpaceBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSpaceBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSpaceBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /spaces/{space_id}][%d] updateSpaceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpaceNotFound creates a UpdateSpaceNotFound with default headers values
func NewUpdateSpaceNotFound() *UpdateSpaceNotFound {
	return &UpdateSpaceNotFound{}
}

/*UpdateSpaceNotFound handles this case with default header values.

Not Found
*/
type UpdateSpaceNotFound struct {
	Payload *models.Error
}

func (o *UpdateSpaceNotFound) Error() string {
	return fmt.Sprintf("[PATCH /spaces/{space_id}][%d] updateSpaceNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSpaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSpaceUnprocessableEntity creates a UpdateSpaceUnprocessableEntity with default headers values
func NewUpdateSpaceUnprocessableEntity() *UpdateSpaceUnprocessableEntity {
	return &UpdateSpaceUnprocessableEntity{}
}

/*UpdateSpaceUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateSpaceUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateSpaceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /spaces/{space_id}][%d] updateSpaceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateSpaceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}