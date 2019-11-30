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

// ScheduledPlansForLookReader is a Reader for the ScheduledPlansForLook structure.
type ScheduledPlansForLookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduledPlansForLookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduledPlansForLookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScheduledPlansForLookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScheduledPlansForLookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewScheduledPlansForLookOK creates a ScheduledPlansForLookOK with default headers values
func NewScheduledPlansForLookOK() *ScheduledPlansForLookOK {
	return &ScheduledPlansForLookOK{}
}

/*ScheduledPlansForLookOK handles this case with default header values.

Scheduled Plan
*/
type ScheduledPlansForLookOK struct {
	Payload []*models.ScheduledPlan
}

func (o *ScheduledPlansForLookOK) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/look/{look_id}][%d] scheduledPlansForLookOK  %+v", 200, o.Payload)
}

func (o *ScheduledPlansForLookOK) GetPayload() []*models.ScheduledPlan {
	return o.Payload
}

func (o *ScheduledPlansForLookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForLookBadRequest creates a ScheduledPlansForLookBadRequest with default headers values
func NewScheduledPlansForLookBadRequest() *ScheduledPlansForLookBadRequest {
	return &ScheduledPlansForLookBadRequest{}
}

/*ScheduledPlansForLookBadRequest handles this case with default header values.

Bad Request
*/
type ScheduledPlansForLookBadRequest struct {
	Payload *models.Error
}

func (o *ScheduledPlansForLookBadRequest) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/look/{look_id}][%d] scheduledPlansForLookBadRequest  %+v", 400, o.Payload)
}

func (o *ScheduledPlansForLookBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ScheduledPlansForLookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForLookNotFound creates a ScheduledPlansForLookNotFound with default headers values
func NewScheduledPlansForLookNotFound() *ScheduledPlansForLookNotFound {
	return &ScheduledPlansForLookNotFound{}
}

/*ScheduledPlansForLookNotFound handles this case with default header values.

Not Found
*/
type ScheduledPlansForLookNotFound struct {
	Payload *models.Error
}

func (o *ScheduledPlansForLookNotFound) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/look/{look_id}][%d] scheduledPlansForLookNotFound  %+v", 404, o.Payload)
}

func (o *ScheduledPlansForLookNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ScheduledPlansForLookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
