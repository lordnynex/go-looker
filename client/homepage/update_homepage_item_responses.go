// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdateHomepageItemReader is a Reader for the UpdateHomepageItem structure.
type UpdateHomepageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHomepageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateHomepageItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateHomepageItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateHomepageItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateHomepageItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateHomepageItemOK creates a UpdateHomepageItemOK with default headers values
func NewUpdateHomepageItemOK() *UpdateHomepageItemOK {
	return &UpdateHomepageItemOK{}
}

/*UpdateHomepageItemOK handles this case with default header values.

Homepage Item
*/
type UpdateHomepageItemOK struct {
	Payload *models.HomepageItem
}

func (o *UpdateHomepageItemOK) Error() string {
	return fmt.Sprintf("[PATCH /homepage_items/{homepage_item_id}][%d] updateHomepageItemOK  %+v", 200, o.Payload)
}

func (o *UpdateHomepageItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomepageItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageItemBadRequest creates a UpdateHomepageItemBadRequest with default headers values
func NewUpdateHomepageItemBadRequest() *UpdateHomepageItemBadRequest {
	return &UpdateHomepageItemBadRequest{}
}

/*UpdateHomepageItemBadRequest handles this case with default header values.

Bad Request
*/
type UpdateHomepageItemBadRequest struct {
	Payload *models.Error
}

func (o *UpdateHomepageItemBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /homepage_items/{homepage_item_id}][%d] updateHomepageItemBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateHomepageItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageItemNotFound creates a UpdateHomepageItemNotFound with default headers values
func NewUpdateHomepageItemNotFound() *UpdateHomepageItemNotFound {
	return &UpdateHomepageItemNotFound{}
}

/*UpdateHomepageItemNotFound handles this case with default header values.

Not Found
*/
type UpdateHomepageItemNotFound struct {
	Payload *models.Error
}

func (o *UpdateHomepageItemNotFound) Error() string {
	return fmt.Sprintf("[PATCH /homepage_items/{homepage_item_id}][%d] updateHomepageItemNotFound  %+v", 404, o.Payload)
}

func (o *UpdateHomepageItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHomepageItemUnprocessableEntity creates a UpdateHomepageItemUnprocessableEntity with default headers values
func NewUpdateHomepageItemUnprocessableEntity() *UpdateHomepageItemUnprocessableEntity {
	return &UpdateHomepageItemUnprocessableEntity{}
}

/*UpdateHomepageItemUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateHomepageItemUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateHomepageItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /homepage_items/{homepage_item_id}][%d] updateHomepageItemUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateHomepageItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
