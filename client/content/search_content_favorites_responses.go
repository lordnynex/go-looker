// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// SearchContentFavoritesReader is a Reader for the SearchContentFavorites structure.
type SearchContentFavoritesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchContentFavoritesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchContentFavoritesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSearchContentFavoritesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSearchContentFavoritesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchContentFavoritesOK creates a SearchContentFavoritesOK with default headers values
func NewSearchContentFavoritesOK() *SearchContentFavoritesOK {
	return &SearchContentFavoritesOK{}
}

/*SearchContentFavoritesOK handles this case with default header values.

Favorite Content
*/
type SearchContentFavoritesOK struct {
	Payload []*models.ContentFavorite
}

func (o *SearchContentFavoritesOK) Error() string {
	return fmt.Sprintf("[GET /content_favorite/search][%d] searchContentFavoritesOK  %+v", 200, o.Payload)
}

func (o *SearchContentFavoritesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchContentFavoritesBadRequest creates a SearchContentFavoritesBadRequest with default headers values
func NewSearchContentFavoritesBadRequest() *SearchContentFavoritesBadRequest {
	return &SearchContentFavoritesBadRequest{}
}

/*SearchContentFavoritesBadRequest handles this case with default header values.

Bad Request
*/
type SearchContentFavoritesBadRequest struct {
	Payload *models.Error
}

func (o *SearchContentFavoritesBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_favorite/search][%d] searchContentFavoritesBadRequest  %+v", 400, o.Payload)
}

func (o *SearchContentFavoritesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchContentFavoritesNotFound creates a SearchContentFavoritesNotFound with default headers values
func NewSearchContentFavoritesNotFound() *SearchContentFavoritesNotFound {
	return &SearchContentFavoritesNotFound{}
}

/*SearchContentFavoritesNotFound handles this case with default header values.

Not Found
*/
type SearchContentFavoritesNotFound struct {
	Payload *models.Error
}

func (o *SearchContentFavoritesNotFound) Error() string {
	return fmt.Sprintf("[GET /content_favorite/search][%d] searchContentFavoritesNotFound  %+v", 404, o.Payload)
}

func (o *SearchContentFavoritesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
