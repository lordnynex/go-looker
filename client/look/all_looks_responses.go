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

// AllLooksReader is a Reader for the AllLooks structure.
type AllLooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllLooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllLooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllLooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllLooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllLooksOK creates a AllLooksOK with default headers values
func NewAllLooksOK() *AllLooksOK {
	return &AllLooksOK{}
}

/*AllLooksOK handles this case with default header values.

Look
*/
type AllLooksOK struct {
	Payload []*models.Look
}

func (o *AllLooksOK) Error() string {
	return fmt.Sprintf("[GET /looks][%d] allLooksOK  %+v", 200, o.Payload)
}

func (o *AllLooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllLooksBadRequest creates a AllLooksBadRequest with default headers values
func NewAllLooksBadRequest() *AllLooksBadRequest {
	return &AllLooksBadRequest{}
}

/*AllLooksBadRequest handles this case with default header values.

Bad Request
*/
type AllLooksBadRequest struct {
	Payload *models.Error
}

func (o *AllLooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /looks][%d] allLooksBadRequest  %+v", 400, o.Payload)
}

func (o *AllLooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllLooksNotFound creates a AllLooksNotFound with default headers values
func NewAllLooksNotFound() *AllLooksNotFound {
	return &AllLooksNotFound{}
}

/*AllLooksNotFound handles this case with default header values.

Not Found
*/
type AllLooksNotFound struct {
	Payload *models.Error
}

func (o *AllLooksNotFound) Error() string {
	return fmt.Sprintf("[GET /looks][%d] allLooksNotFound  %+v", 404, o.Payload)
}

func (o *AllLooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
