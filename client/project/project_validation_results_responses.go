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

// ProjectValidationResultsReader is a Reader for the ProjectValidationResults structure.
type ProjectValidationResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectValidationResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectValidationResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewProjectValidationResultsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProjectValidationResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectValidationResultsOK creates a ProjectValidationResultsOK with default headers values
func NewProjectValidationResultsOK() *ProjectValidationResultsOK {
	return &ProjectValidationResultsOK{}
}

/*ProjectValidationResultsOK handles this case with default header values.

Project validation results
*/
type ProjectValidationResultsOK struct {
	Payload *models.ProjectValidationCache
}

func (o *ProjectValidationResultsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/validate][%d] projectValidationResultsOK  %+v", 200, o.Payload)
}

func (o *ProjectValidationResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectValidationCache)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectValidationResultsNoContent creates a ProjectValidationResultsNoContent with default headers values
func NewProjectValidationResultsNoContent() *ProjectValidationResultsNoContent {
	return &ProjectValidationResultsNoContent{}
}

/*ProjectValidationResultsNoContent handles this case with default header values.

Deleted
*/
type ProjectValidationResultsNoContent struct {
}

func (o *ProjectValidationResultsNoContent) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/validate][%d] projectValidationResultsNoContent ", 204)
}

func (o *ProjectValidationResultsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectValidationResultsNotFound creates a ProjectValidationResultsNotFound with default headers values
func NewProjectValidationResultsNotFound() *ProjectValidationResultsNotFound {
	return &ProjectValidationResultsNotFound{}
}

/*ProjectValidationResultsNotFound handles this case with default header values.

Not Found
*/
type ProjectValidationResultsNotFound struct {
	Payload *models.Error
}

func (o *ProjectValidationResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/validate][%d] projectValidationResultsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectValidationResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
