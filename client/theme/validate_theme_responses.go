// Code generated by go-swagger; DO NOT EDIT.

package theme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// ValidateThemeReader is a Reader for the ValidateTheme structure.
type ValidateThemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateThemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateThemeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateThemeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateThemeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewValidateThemeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateThemeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateThemeOK creates a ValidateThemeOK with default headers values
func NewValidateThemeOK() *ValidateThemeOK {
	return &ValidateThemeOK{}
}

/*ValidateThemeOK handles this case with default header values.

Theme validation results
*/
type ValidateThemeOK struct {
	Payload *models.ValidationError
}

func (o *ValidateThemeOK) Error() string {
	return fmt.Sprintf("[POST /themes/validate][%d] validateThemeOK  %+v", 200, o.Payload)
}

func (o *ValidateThemeOK) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ValidateThemeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateThemeBadRequest creates a ValidateThemeBadRequest with default headers values
func NewValidateThemeBadRequest() *ValidateThemeBadRequest {
	return &ValidateThemeBadRequest{}
}

/*ValidateThemeBadRequest handles this case with default header values.

Bad Request
*/
type ValidateThemeBadRequest struct {
	Payload *models.Error
}

func (o *ValidateThemeBadRequest) Error() string {
	return fmt.Sprintf("[POST /themes/validate][%d] validateThemeBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateThemeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateThemeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateThemeNotFound creates a ValidateThemeNotFound with default headers values
func NewValidateThemeNotFound() *ValidateThemeNotFound {
	return &ValidateThemeNotFound{}
}

/*ValidateThemeNotFound handles this case with default header values.

Not Found
*/
type ValidateThemeNotFound struct {
	Payload *models.Error
}

func (o *ValidateThemeNotFound) Error() string {
	return fmt.Sprintf("[POST /themes/validate][%d] validateThemeNotFound  %+v", 404, o.Payload)
}

func (o *ValidateThemeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateThemeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateThemeConflict creates a ValidateThemeConflict with default headers values
func NewValidateThemeConflict() *ValidateThemeConflict {
	return &ValidateThemeConflict{}
}

/*ValidateThemeConflict handles this case with default header values.

Resource Already Exists
*/
type ValidateThemeConflict struct {
	Payload *models.Error
}

func (o *ValidateThemeConflict) Error() string {
	return fmt.Sprintf("[POST /themes/validate][%d] validateThemeConflict  %+v", 409, o.Payload)
}

func (o *ValidateThemeConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateThemeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateThemeUnprocessableEntity creates a ValidateThemeUnprocessableEntity with default headers values
func NewValidateThemeUnprocessableEntity() *ValidateThemeUnprocessableEntity {
	return &ValidateThemeUnprocessableEntity{}
}

/*ValidateThemeUnprocessableEntity handles this case with default header values.

Validation Error
*/
type ValidateThemeUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ValidateThemeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /themes/validate][%d] validateThemeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ValidateThemeUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ValidateThemeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
