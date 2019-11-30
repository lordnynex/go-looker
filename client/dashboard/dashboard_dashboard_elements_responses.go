// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// DashboardDashboardElementsReader is a Reader for the DashboardDashboardElements structure.
type DashboardDashboardElementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardDashboardElementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardDashboardElementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardDashboardElementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardDashboardElementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDashboardDashboardElementsOK creates a DashboardDashboardElementsOK with default headers values
func NewDashboardDashboardElementsOK() *DashboardDashboardElementsOK {
	return &DashboardDashboardElementsOK{}
}

/*DashboardDashboardElementsOK handles this case with default header values.

DashboardElement
*/
type DashboardDashboardElementsOK struct {
	Payload []*models.DashboardElement
}

func (o *DashboardDashboardElementsOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_elements][%d] dashboardDashboardElementsOK  %+v", 200, o.Payload)
}

func (o *DashboardDashboardElementsOK) GetPayload() []*models.DashboardElement {
	return o.Payload
}

func (o *DashboardDashboardElementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardElementsBadRequest creates a DashboardDashboardElementsBadRequest with default headers values
func NewDashboardDashboardElementsBadRequest() *DashboardDashboardElementsBadRequest {
	return &DashboardDashboardElementsBadRequest{}
}

/*DashboardDashboardElementsBadRequest handles this case with default header values.

Bad Request
*/
type DashboardDashboardElementsBadRequest struct {
	Payload *models.Error
}

func (o *DashboardDashboardElementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_elements][%d] dashboardDashboardElementsBadRequest  %+v", 400, o.Payload)
}

func (o *DashboardDashboardElementsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardDashboardElementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardElementsNotFound creates a DashboardDashboardElementsNotFound with default headers values
func NewDashboardDashboardElementsNotFound() *DashboardDashboardElementsNotFound {
	return &DashboardDashboardElementsNotFound{}
}

/*DashboardDashboardElementsNotFound handles this case with default header values.

Not Found
*/
type DashboardDashboardElementsNotFound struct {
	Payload *models.Error
}

func (o *DashboardDashboardElementsNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_elements][%d] dashboardDashboardElementsNotFound  %+v", 404, o.Payload)
}

func (o *DashboardDashboardElementsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DashboardDashboardElementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
