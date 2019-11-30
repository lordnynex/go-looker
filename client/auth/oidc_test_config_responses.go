// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// OidcTestConfigReader is a Reader for the OidcTestConfig structure.
type OidcTestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcTestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOidcTestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOidcTestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOidcTestConfigOK creates a OidcTestConfigOK with default headers values
func NewOidcTestConfigOK() *OidcTestConfigOK {
	return &OidcTestConfigOK{}
}

/*OidcTestConfigOK handles this case with default header values.

OIDC test config.
*/
type OidcTestConfigOK struct {
	Payload *models.OIDCConfig
}

func (o *OidcTestConfigOK) Error() string {
	return fmt.Sprintf("[GET /oidc_test_configs/{test_slug}][%d] oidcTestConfigOK  %+v", 200, o.Payload)
}

func (o *OidcTestConfigOK) GetPayload() *models.OIDCConfig {
	return o.Payload
}

func (o *OidcTestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OIDCConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOidcTestConfigNotFound creates a OidcTestConfigNotFound with default headers values
func NewOidcTestConfigNotFound() *OidcTestConfigNotFound {
	return &OidcTestConfigNotFound{}
}

/*OidcTestConfigNotFound handles this case with default header values.

Not Found
*/
type OidcTestConfigNotFound struct {
	Payload *models.Error
}

func (o *OidcTestConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /oidc_test_configs/{test_slug}][%d] oidcTestConfigNotFound  %+v", 404, o.Payload)
}

func (o *OidcTestConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *OidcTestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
