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

// LookmlModelReader is a Reader for the LookmlModel structure.
type LookmlModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookmlModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLookmlModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLookmlModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLookmlModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLookmlModelOK creates a LookmlModelOK with default headers values
func NewLookmlModelOK() *LookmlModelOK {
	return &LookmlModelOK{}
}

/*LookmlModelOK handles this case with default header values.

LookML Model
*/
type LookmlModelOK struct {
	Payload *models.LookmlModel
}

func (o *LookmlModelOK) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}][%d] lookmlModelOK  %+v", 200, o.Payload)
}

func (o *LookmlModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookmlModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookmlModelBadRequest creates a LookmlModelBadRequest with default headers values
func NewLookmlModelBadRequest() *LookmlModelBadRequest {
	return &LookmlModelBadRequest{}
}

/*LookmlModelBadRequest handles this case with default header values.

Bad Request
*/
type LookmlModelBadRequest struct {
	Payload *models.Error
}

func (o *LookmlModelBadRequest) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}][%d] lookmlModelBadRequest  %+v", 400, o.Payload)
}

func (o *LookmlModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookmlModelNotFound creates a LookmlModelNotFound with default headers values
func NewLookmlModelNotFound() *LookmlModelNotFound {
	return &LookmlModelNotFound{}
}

/*LookmlModelNotFound handles this case with default header values.

Not Found
*/
type LookmlModelNotFound struct {
	Payload *models.Error
}

func (o *LookmlModelNotFound) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}][%d] lookmlModelNotFound  %+v", 404, o.Payload)
}

func (o *LookmlModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
