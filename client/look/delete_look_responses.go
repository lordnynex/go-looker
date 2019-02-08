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

// DeleteLookReader is a Reader for the DeleteLook structure.
type DeleteLookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLookNoContent creates a DeleteLookNoContent with default headers values
func NewDeleteLookNoContent() *DeleteLookNoContent {
	return &DeleteLookNoContent{}
}

/*DeleteLookNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteLookNoContent struct {
	Payload string
}

func (o *DeleteLookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /looks/{look_id}][%d] deleteLookNoContent  %+v", 204, o.Payload)
}

func (o *DeleteLookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLookBadRequest creates a DeleteLookBadRequest with default headers values
func NewDeleteLookBadRequest() *DeleteLookBadRequest {
	return &DeleteLookBadRequest{}
}

/*DeleteLookBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLookBadRequest struct {
	Payload *models.Error
}

func (o *DeleteLookBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /looks/{look_id}][%d] deleteLookBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLookNotFound creates a DeleteLookNotFound with default headers values
func NewDeleteLookNotFound() *DeleteLookNotFound {
	return &DeleteLookNotFound{}
}

/*DeleteLookNotFound handles this case with default header values.

Not Found
*/
type DeleteLookNotFound struct {
	Payload *models.Error
}

func (o *DeleteLookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /looks/{look_id}][%d] deleteLookNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
