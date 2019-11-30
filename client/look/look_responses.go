// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// LookReader is a Reader for the Look structure.
type LookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLookOK creates a LookOK with default headers values
func NewLookOK() *LookOK {
	return &LookOK{}
}

/*LookOK handles this case with default header values.

Look
*/
type LookOK struct {
	Payload *models.LookWithQuery
}

func (o *LookOK) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}][%d] lookOK  %+v", 200, o.Payload)
}

func (o *LookOK) GetPayload() *models.LookWithQuery {
	return o.Payload
}

func (o *LookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookWithQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookBadRequest creates a LookBadRequest with default headers values
func NewLookBadRequest() *LookBadRequest {
	return &LookBadRequest{}
}

/*LookBadRequest handles this case with default header values.

Bad Request
*/
type LookBadRequest struct {
	Payload *models.Error
}

func (o *LookBadRequest) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}][%d] lookBadRequest  %+v", 400, o.Payload)
}

func (o *LookBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *LookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookNotFound creates a LookNotFound with default headers values
func NewLookNotFound() *LookNotFound {
	return &LookNotFound{}
}

/*LookNotFound handles this case with default header values.

Not Found
*/
type LookNotFound struct {
	Payload *models.Error
}

func (o *LookNotFound) Error() string {
	return fmt.Sprintf("[GET /looks/{look_id}][%d] lookNotFound  %+v", 404, o.Payload)
}

func (o *LookNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *LookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
