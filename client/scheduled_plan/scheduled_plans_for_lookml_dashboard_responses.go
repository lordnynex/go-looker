// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// ScheduledPlansForLookmlDashboardReader is a Reader for the ScheduledPlansForLookmlDashboard structure.
type ScheduledPlansForLookmlDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduledPlansForLookmlDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewScheduledPlansForLookmlDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewScheduledPlansForLookmlDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewScheduledPlansForLookmlDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewScheduledPlansForLookmlDashboardOK creates a ScheduledPlansForLookmlDashboardOK with default headers values
func NewScheduledPlansForLookmlDashboardOK() *ScheduledPlansForLookmlDashboardOK {
	return &ScheduledPlansForLookmlDashboardOK{}
}

/*ScheduledPlansForLookmlDashboardOK handles this case with default header values.

Scheduled Plan
*/
type ScheduledPlansForLookmlDashboardOK struct {
	Payload []*models.ScheduledPlan
}

func (o *ScheduledPlansForLookmlDashboardOK) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/lookml_dashboard/{lookml_dashboard_id}][%d] scheduledPlansForLookmlDashboardOK  %+v", 200, o.Payload)
}

func (o *ScheduledPlansForLookmlDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForLookmlDashboardBadRequest creates a ScheduledPlansForLookmlDashboardBadRequest with default headers values
func NewScheduledPlansForLookmlDashboardBadRequest() *ScheduledPlansForLookmlDashboardBadRequest {
	return &ScheduledPlansForLookmlDashboardBadRequest{}
}

/*ScheduledPlansForLookmlDashboardBadRequest handles this case with default header values.

Bad Request
*/
type ScheduledPlansForLookmlDashboardBadRequest struct {
	Payload *models.Error
}

func (o *ScheduledPlansForLookmlDashboardBadRequest) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/lookml_dashboard/{lookml_dashboard_id}][%d] scheduledPlansForLookmlDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *ScheduledPlansForLookmlDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForLookmlDashboardNotFound creates a ScheduledPlansForLookmlDashboardNotFound with default headers values
func NewScheduledPlansForLookmlDashboardNotFound() *ScheduledPlansForLookmlDashboardNotFound {
	return &ScheduledPlansForLookmlDashboardNotFound{}
}

/*ScheduledPlansForLookmlDashboardNotFound handles this case with default header values.

Not Found
*/
type ScheduledPlansForLookmlDashboardNotFound struct {
	Payload *models.Error
}

func (o *ScheduledPlansForLookmlDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/lookml_dashboard/{lookml_dashboard_id}][%d] scheduledPlansForLookmlDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ScheduledPlansForLookmlDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
