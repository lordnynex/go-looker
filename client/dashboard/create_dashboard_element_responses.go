// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// CreateDashboardElementReader is a Reader for the CreateDashboardElement structure.
type CreateDashboardElementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardElementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDashboardElementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDashboardElementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDashboardElementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDashboardElementConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateDashboardElementUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardElementOK creates a CreateDashboardElementOK with default headers values
func NewCreateDashboardElementOK() *CreateDashboardElementOK {
	return &CreateDashboardElementOK{}
}

/*CreateDashboardElementOK handles this case with default header values.

DashboardElement
*/
type CreateDashboardElementOK struct {
	Payload *models.DashboardElement
}

func (o *CreateDashboardElementOK) Error() string {
	return fmt.Sprintf("[POST /dashboard_elements][%d] createDashboardElementOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardElementOK) GetPayload() *models.DashboardElement {
	return o.Payload
}

func (o *CreateDashboardElementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardElement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardElementBadRequest creates a CreateDashboardElementBadRequest with default headers values
func NewCreateDashboardElementBadRequest() *CreateDashboardElementBadRequest {
	return &CreateDashboardElementBadRequest{}
}

/*CreateDashboardElementBadRequest handles this case with default header values.

Bad Request
*/
type CreateDashboardElementBadRequest struct {
	Payload *models.Error
}

func (o *CreateDashboardElementBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboard_elements][%d] createDashboardElementBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDashboardElementBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardElementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardElementNotFound creates a CreateDashboardElementNotFound with default headers values
func NewCreateDashboardElementNotFound() *CreateDashboardElementNotFound {
	return &CreateDashboardElementNotFound{}
}

/*CreateDashboardElementNotFound handles this case with default header values.

Not Found
*/
type CreateDashboardElementNotFound struct {
	Payload *models.Error
}

func (o *CreateDashboardElementNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboard_elements][%d] createDashboardElementNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardElementNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardElementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardElementConflict creates a CreateDashboardElementConflict with default headers values
func NewCreateDashboardElementConflict() *CreateDashboardElementConflict {
	return &CreateDashboardElementConflict{}
}

/*CreateDashboardElementConflict handles this case with default header values.

Resource Already Exists
*/
type CreateDashboardElementConflict struct {
	Payload *models.Error
}

func (o *CreateDashboardElementConflict) Error() string {
	return fmt.Sprintf("[POST /dashboard_elements][%d] createDashboardElementConflict  %+v", 409, o.Payload)
}

func (o *CreateDashboardElementConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardElementConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardElementUnprocessableEntity creates a CreateDashboardElementUnprocessableEntity with default headers values
func NewCreateDashboardElementUnprocessableEntity() *CreateDashboardElementUnprocessableEntity {
	return &CreateDashboardElementUnprocessableEntity{}
}

/*CreateDashboardElementUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateDashboardElementUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateDashboardElementUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboard_elements][%d] createDashboardElementUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateDashboardElementUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateDashboardElementUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
