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

// UpdateDashboardReader is a Reader for the UpdateDashboard structure.
type UpdateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateDashboardMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDashboardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDashboardOK creates a UpdateDashboardOK with default headers values
func NewUpdateDashboardOK() *UpdateDashboardOK {
	return &UpdateDashboardOK{}
}

/*UpdateDashboardOK handles this case with default header values.

Dashboard
*/
type UpdateDashboardOK struct {
	Payload *models.Dashboard
}

func (o *UpdateDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardOK) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *UpdateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardBadRequest creates a UpdateDashboardBadRequest with default headers values
func NewUpdateDashboardBadRequest() *UpdateDashboardBadRequest {
	return &UpdateDashboardBadRequest{}
}

/*UpdateDashboardBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDashboardBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardNotFound creates a UpdateDashboardNotFound with default headers values
func NewUpdateDashboardNotFound() *UpdateDashboardNotFound {
	return &UpdateDashboardNotFound{}
}

/*UpdateDashboardNotFound handles this case with default header values.

Not Found
*/
type UpdateDashboardNotFound struct {
	Payload *models.Error
}

func (o *UpdateDashboardNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardMethodNotAllowed creates a UpdateDashboardMethodNotAllowed with default headers values
func NewUpdateDashboardMethodNotAllowed() *UpdateDashboardMethodNotAllowed {
	return &UpdateDashboardMethodNotAllowed{}
}

/*UpdateDashboardMethodNotAllowed handles this case with default header values.

Resource Can't Be Modified
*/
type UpdateDashboardMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateDashboardMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateDashboardMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardUnprocessableEntity creates a UpdateDashboardUnprocessableEntity with default headers values
func NewUpdateDashboardUnprocessableEntity() *UpdateDashboardUnprocessableEntity {
	return &UpdateDashboardUnprocessableEntity{}
}

/*UpdateDashboardUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateDashboardUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateDashboardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/{dashboard_id}][%d] updateDashboardUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateDashboardUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateDashboardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
