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

// DeleteDashboardLayoutReader is a Reader for the DeleteDashboardLayout structure.
type DeleteDashboardLayoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardLayoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDashboardLayoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteDashboardLayoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteDashboardLayoutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDashboardLayoutNoContent creates a DeleteDashboardLayoutNoContent with default headers values
func NewDeleteDashboardLayoutNoContent() *DeleteDashboardLayoutNoContent {
	return &DeleteDashboardLayoutNoContent{}
}

/*DeleteDashboardLayoutNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteDashboardLayoutNoContent struct {
	Payload string
}

func (o *DeleteDashboardLayoutNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dashboard_layouts/{dashboard_layout_id}][%d] deleteDashboardLayoutNoContent  %+v", 204, o.Payload)
}

func (o *DeleteDashboardLayoutNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteDashboardLayoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardLayoutNotFound creates a DeleteDashboardLayoutNotFound with default headers values
func NewDeleteDashboardLayoutNotFound() *DeleteDashboardLayoutNotFound {
	return &DeleteDashboardLayoutNotFound{}
}

/*DeleteDashboardLayoutNotFound handles this case with default header values.

Not Found
*/
type DeleteDashboardLayoutNotFound struct {
	Payload *models.Error
}

func (o *DeleteDashboardLayoutNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dashboard_layouts/{dashboard_layout_id}][%d] deleteDashboardLayoutNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDashboardLayoutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDashboardLayoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardLayoutUnprocessableEntity creates a DeleteDashboardLayoutUnprocessableEntity with default headers values
func NewDeleteDashboardLayoutUnprocessableEntity() *DeleteDashboardLayoutUnprocessableEntity {
	return &DeleteDashboardLayoutUnprocessableEntity{}
}

/*DeleteDashboardLayoutUnprocessableEntity handles this case with default header values.

Validation Error
*/
type DeleteDashboardLayoutUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *DeleteDashboardLayoutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /dashboard_layouts/{dashboard_layout_id}][%d] deleteDashboardLayoutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteDashboardLayoutUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *DeleteDashboardLayoutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
