// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*UpdateProjectOK handles this case with default header values.

Project
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectBadRequest creates a UpdateProjectBadRequest with default headers values
func NewUpdateProjectBadRequest() *UpdateProjectBadRequest {
	return &UpdateProjectBadRequest{}
}

/*UpdateProjectBadRequest handles this case with default header values.

Bad Request
*/
type UpdateProjectBadRequest struct {
	Payload *models.Error
}

func (o *UpdateProjectBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/*UpdateProjectNotFound handles this case with default header values.

Not Found
*/
type UpdateProjectNotFound struct {
	Payload *models.Error
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectConflict creates a UpdateProjectConflict with default headers values
func NewUpdateProjectConflict() *UpdateProjectConflict {
	return &UpdateProjectConflict{}
}

/*UpdateProjectConflict handles this case with default header values.

Resource Already Exists
*/
type UpdateProjectConflict struct {
	Payload *models.Error
}

func (o *UpdateProjectConflict) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectUnprocessableEntity creates a UpdateProjectUnprocessableEntity with default headers values
func NewUpdateProjectUnprocessableEntity() *UpdateProjectUnprocessableEntity {
	return &UpdateProjectUnprocessableEntity{}
}

/*UpdateProjectUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateProjectUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectInternalServerError creates a UpdateProjectInternalServerError with default headers values
func NewUpdateProjectInternalServerError() *UpdateProjectInternalServerError {
	return &UpdateProjectInternalServerError{}
}

/*UpdateProjectInternalServerError handles this case with default header values.

Server Error
*/
type UpdateProjectInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id}][%d] updateProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
