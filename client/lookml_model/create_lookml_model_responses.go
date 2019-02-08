// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// CreateLookmlModelReader is a Reader for the CreateLookmlModel structure.
type CreateLookmlModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLookmlModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLookmlModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLookmlModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLookmlModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateLookmlModelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateLookmlModelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLookmlModelOK creates a CreateLookmlModelOK with default headers values
func NewCreateLookmlModelOK() *CreateLookmlModelOK {
	return &CreateLookmlModelOK{}
}

/*CreateLookmlModelOK handles this case with default header values.

LookML Model
*/
type CreateLookmlModelOK struct {
	Payload *models.LookmlModel
}

func (o *CreateLookmlModelOK) Error() string {
	return fmt.Sprintf("[POST /lookml_models][%d] createLookmlModelOK  %+v", 200, o.Payload)
}

func (o *CreateLookmlModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookmlModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlModelBadRequest creates a CreateLookmlModelBadRequest with default headers values
func NewCreateLookmlModelBadRequest() *CreateLookmlModelBadRequest {
	return &CreateLookmlModelBadRequest{}
}

/*CreateLookmlModelBadRequest handles this case with default header values.

Bad Request
*/
type CreateLookmlModelBadRequest struct {
	Payload *models.Error
}

func (o *CreateLookmlModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /lookml_models][%d] createLookmlModelBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLookmlModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlModelNotFound creates a CreateLookmlModelNotFound with default headers values
func NewCreateLookmlModelNotFound() *CreateLookmlModelNotFound {
	return &CreateLookmlModelNotFound{}
}

/*CreateLookmlModelNotFound handles this case with default header values.

Not Found
*/
type CreateLookmlModelNotFound struct {
	Payload *models.Error
}

func (o *CreateLookmlModelNotFound) Error() string {
	return fmt.Sprintf("[POST /lookml_models][%d] createLookmlModelNotFound  %+v", 404, o.Payload)
}

func (o *CreateLookmlModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlModelConflict creates a CreateLookmlModelConflict with default headers values
func NewCreateLookmlModelConflict() *CreateLookmlModelConflict {
	return &CreateLookmlModelConflict{}
}

/*CreateLookmlModelConflict handles this case with default header values.

Resource Already Exists
*/
type CreateLookmlModelConflict struct {
	Payload *models.Error
}

func (o *CreateLookmlModelConflict) Error() string {
	return fmt.Sprintf("[POST /lookml_models][%d] createLookmlModelConflict  %+v", 409, o.Payload)
}

func (o *CreateLookmlModelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLookmlModelUnprocessableEntity creates a CreateLookmlModelUnprocessableEntity with default headers values
func NewCreateLookmlModelUnprocessableEntity() *CreateLookmlModelUnprocessableEntity {
	return &CreateLookmlModelUnprocessableEntity{}
}

/*CreateLookmlModelUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateLookmlModelUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateLookmlModelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /lookml_models][%d] createLookmlModelUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateLookmlModelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
