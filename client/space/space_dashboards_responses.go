// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// SpaceDashboardsReader is a Reader for the SpaceDashboards structure.
type SpaceDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpaceDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpaceDashboardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpaceDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSpaceDashboardsOK creates a SpaceDashboardsOK with default headers values
func NewSpaceDashboardsOK() *SpaceDashboardsOK {
	return &SpaceDashboardsOK{}
}

/*SpaceDashboardsOK handles this case with default header values.

Dashboard
*/
type SpaceDashboardsOK struct {
	Payload []*models.Dashboard
}

func (o *SpaceDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/dashboards][%d] spaceDashboardsOK  %+v", 200, o.Payload)
}

func (o *SpaceDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceDashboardsBadRequest creates a SpaceDashboardsBadRequest with default headers values
func NewSpaceDashboardsBadRequest() *SpaceDashboardsBadRequest {
	return &SpaceDashboardsBadRequest{}
}

/*SpaceDashboardsBadRequest handles this case with default header values.

Bad Request
*/
type SpaceDashboardsBadRequest struct {
	Payload *models.Error
}

func (o *SpaceDashboardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/dashboards][%d] spaceDashboardsBadRequest  %+v", 400, o.Payload)
}

func (o *SpaceDashboardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceDashboardsNotFound creates a SpaceDashboardsNotFound with default headers values
func NewSpaceDashboardsNotFound() *SpaceDashboardsNotFound {
	return &SpaceDashboardsNotFound{}
}

/*SpaceDashboardsNotFound handles this case with default header values.

Not Found
*/
type SpaceDashboardsNotFound struct {
	Payload *models.Error
}

func (o *SpaceDashboardsNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/dashboards][%d] spaceDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *SpaceDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}