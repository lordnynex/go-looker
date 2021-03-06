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

// HomepageSectionReader is a Reader for the HomepageSection structure.
type HomepageSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HomepageSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHomepageSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHomepageSectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHomepageSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHomepageSectionOK creates a HomepageSectionOK with default headers values
func NewHomepageSectionOK() *HomepageSectionOK {
	return &HomepageSectionOK{}
}

/*HomepageSectionOK handles this case with default header values.

Homepage section
*/
type HomepageSectionOK struct {
	Payload *models.HomepageSection
}

func (o *HomepageSectionOK) Error() string {
	return fmt.Sprintf("[GET /homepage_sections/{homepage_section_id}][%d] homepageSectionOK  %+v", 200, o.Payload)
}

func (o *HomepageSectionOK) GetPayload() *models.HomepageSection {
	return o.Payload
}

func (o *HomepageSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HomepageSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageSectionBadRequest creates a HomepageSectionBadRequest with default headers values
func NewHomepageSectionBadRequest() *HomepageSectionBadRequest {
	return &HomepageSectionBadRequest{}
}

/*HomepageSectionBadRequest handles this case with default header values.

Bad Request
*/
type HomepageSectionBadRequest struct {
	Payload *models.Error
}

func (o *HomepageSectionBadRequest) Error() string {
	return fmt.Sprintf("[GET /homepage_sections/{homepage_section_id}][%d] homepageSectionBadRequest  %+v", 400, o.Payload)
}

func (o *HomepageSectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *HomepageSectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageSectionNotFound creates a HomepageSectionNotFound with default headers values
func NewHomepageSectionNotFound() *HomepageSectionNotFound {
	return &HomepageSectionNotFound{}
}

/*HomepageSectionNotFound handles this case with default header values.

Not Found
*/
type HomepageSectionNotFound struct {
	Payload *models.Error
}

func (o *HomepageSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /homepage_sections/{homepage_section_id}][%d] homepageSectionNotFound  %+v", 404, o.Payload)
}

func (o *HomepageSectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *HomepageSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
