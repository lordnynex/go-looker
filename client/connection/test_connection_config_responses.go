// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// TestConnectionConfigReader is a Reader for the TestConnectionConfig structure.
type TestConnectionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestConnectionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestConnectionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestConnectionConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestConnectionConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewTestConnectionConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestConnectionConfigOK creates a TestConnectionConfigOK with default headers values
func NewTestConnectionConfigOK() *TestConnectionConfigOK {
	return &TestConnectionConfigOK{}
}

/*TestConnectionConfigOK handles this case with default header values.

Test results
*/
type TestConnectionConfigOK struct {
	Payload []*models.DBConnectionTestResult
}

func (o *TestConnectionConfigOK) Error() string {
	return fmt.Sprintf("[PUT /connections/test][%d] testConnectionConfigOK  %+v", 200, o.Payload)
}

func (o *TestConnectionConfigOK) GetPayload() []*models.DBConnectionTestResult {
	return o.Payload
}

func (o *TestConnectionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestConnectionConfigBadRequest creates a TestConnectionConfigBadRequest with default headers values
func NewTestConnectionConfigBadRequest() *TestConnectionConfigBadRequest {
	return &TestConnectionConfigBadRequest{}
}

/*TestConnectionConfigBadRequest handles this case with default header values.

Bad Request
*/
type TestConnectionConfigBadRequest struct {
	Payload *models.Error
}

func (o *TestConnectionConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /connections/test][%d] testConnectionConfigBadRequest  %+v", 400, o.Payload)
}

func (o *TestConnectionConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestConnectionConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestConnectionConfigNotFound creates a TestConnectionConfigNotFound with default headers values
func NewTestConnectionConfigNotFound() *TestConnectionConfigNotFound {
	return &TestConnectionConfigNotFound{}
}

/*TestConnectionConfigNotFound handles this case with default header values.

Not Found
*/
type TestConnectionConfigNotFound struct {
	Payload *models.Error
}

func (o *TestConnectionConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /connections/test][%d] testConnectionConfigNotFound  %+v", 404, o.Payload)
}

func (o *TestConnectionConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestConnectionConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestConnectionConfigTooManyRequests creates a TestConnectionConfigTooManyRequests with default headers values
func NewTestConnectionConfigTooManyRequests() *TestConnectionConfigTooManyRequests {
	return &TestConnectionConfigTooManyRequests{}
}

/*TestConnectionConfigTooManyRequests handles this case with default header values.

Too Many Requests
*/
type TestConnectionConfigTooManyRequests struct {
	Payload *models.Error
}

func (o *TestConnectionConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /connections/test][%d] testConnectionConfigTooManyRequests  %+v", 429, o.Payload)
}

func (o *TestConnectionConfigTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestConnectionConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
