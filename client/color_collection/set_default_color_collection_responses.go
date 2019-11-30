// Code generated by go-swagger; DO NOT EDIT.

package color_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// SetDefaultColorCollectionReader is a Reader for the SetDefaultColorCollection structure.
type SetDefaultColorCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultColorCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDefaultColorCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDefaultColorCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetDefaultColorCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSetDefaultColorCollectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDefaultColorCollectionOK creates a SetDefaultColorCollectionOK with default headers values
func NewSetDefaultColorCollectionOK() *SetDefaultColorCollectionOK {
	return &SetDefaultColorCollectionOK{}
}

/*SetDefaultColorCollectionOK handles this case with default header values.

ColorCollection
*/
type SetDefaultColorCollectionOK struct {
	Payload *models.ColorCollection
}

func (o *SetDefaultColorCollectionOK) Error() string {
	return fmt.Sprintf("[PUT /color_collections/default][%d] setDefaultColorCollectionOK  %+v", 200, o.Payload)
}

func (o *SetDefaultColorCollectionOK) GetPayload() *models.ColorCollection {
	return o.Payload
}

func (o *SetDefaultColorCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ColorCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultColorCollectionBadRequest creates a SetDefaultColorCollectionBadRequest with default headers values
func NewSetDefaultColorCollectionBadRequest() *SetDefaultColorCollectionBadRequest {
	return &SetDefaultColorCollectionBadRequest{}
}

/*SetDefaultColorCollectionBadRequest handles this case with default header values.

Bad Request
*/
type SetDefaultColorCollectionBadRequest struct {
	Payload *models.Error
}

func (o *SetDefaultColorCollectionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /color_collections/default][%d] setDefaultColorCollectionBadRequest  %+v", 400, o.Payload)
}

func (o *SetDefaultColorCollectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultColorCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultColorCollectionNotFound creates a SetDefaultColorCollectionNotFound with default headers values
func NewSetDefaultColorCollectionNotFound() *SetDefaultColorCollectionNotFound {
	return &SetDefaultColorCollectionNotFound{}
}

/*SetDefaultColorCollectionNotFound handles this case with default header values.

Not Found
*/
type SetDefaultColorCollectionNotFound struct {
	Payload *models.Error
}

func (o *SetDefaultColorCollectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /color_collections/default][%d] setDefaultColorCollectionNotFound  %+v", 404, o.Payload)
}

func (o *SetDefaultColorCollectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultColorCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultColorCollectionUnprocessableEntity creates a SetDefaultColorCollectionUnprocessableEntity with default headers values
func NewSetDefaultColorCollectionUnprocessableEntity() *SetDefaultColorCollectionUnprocessableEntity {
	return &SetDefaultColorCollectionUnprocessableEntity{}
}

/*SetDefaultColorCollectionUnprocessableEntity handles this case with default header values.

Validation Error
*/
type SetDefaultColorCollectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *SetDefaultColorCollectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /color_collections/default][%d] setDefaultColorCollectionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SetDefaultColorCollectionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *SetDefaultColorCollectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}