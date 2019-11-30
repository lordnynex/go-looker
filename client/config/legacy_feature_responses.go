// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// LegacyFeatureReader is a Reader for the LegacyFeature structure.
type LegacyFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LegacyFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLegacyFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLegacyFeatureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLegacyFeatureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLegacyFeatureOK creates a LegacyFeatureOK with default headers values
func NewLegacyFeatureOK() *LegacyFeatureOK {
	return &LegacyFeatureOK{}
}

/*LegacyFeatureOK handles this case with default header values.

Legacy Feature
*/
type LegacyFeatureOK struct {
	Payload *models.LegacyFeature
}

func (o *LegacyFeatureOK) Error() string {
	return fmt.Sprintf("[GET /legacy_features/{legacy_feature_id}][%d] legacyFeatureOK  %+v", 200, o.Payload)
}

func (o *LegacyFeatureOK) GetPayload() *models.LegacyFeature {
	return o.Payload
}

func (o *LegacyFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LegacyFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLegacyFeatureBadRequest creates a LegacyFeatureBadRequest with default headers values
func NewLegacyFeatureBadRequest() *LegacyFeatureBadRequest {
	return &LegacyFeatureBadRequest{}
}

/*LegacyFeatureBadRequest handles this case with default header values.

Bad Request
*/
type LegacyFeatureBadRequest struct {
	Payload *models.Error
}

func (o *LegacyFeatureBadRequest) Error() string {
	return fmt.Sprintf("[GET /legacy_features/{legacy_feature_id}][%d] legacyFeatureBadRequest  %+v", 400, o.Payload)
}

func (o *LegacyFeatureBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *LegacyFeatureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLegacyFeatureNotFound creates a LegacyFeatureNotFound with default headers values
func NewLegacyFeatureNotFound() *LegacyFeatureNotFound {
	return &LegacyFeatureNotFound{}
}

/*LegacyFeatureNotFound handles this case with default header values.

Not Found
*/
type LegacyFeatureNotFound struct {
	Payload *models.Error
}

func (o *LegacyFeatureNotFound) Error() string {
	return fmt.Sprintf("[GET /legacy_features/{legacy_feature_id}][%d] legacyFeatureNotFound  %+v", 404, o.Payload)
}

func (o *LegacyFeatureNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *LegacyFeatureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
