package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// GetDhcpOptionsReader is a Reader for the GetDhcpOptions structure.
type GetDhcpOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDhcpOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDhcpOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDhcpOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDhcpOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDhcpOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDhcpOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDhcpOptionsOK creates a GetDhcpOptionsOK with default headers values
func NewGetDhcpOptionsOK() *GetDhcpOptionsOK {
	return &GetDhcpOptionsOK{}
}

/*GetDhcpOptionsOK handles this case with default header values.

The set of DHCP options was retrieved.
*/
type GetDhcpOptionsOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.DhcpOptions
}

func (o *GetDhcpOptionsOK) Error() string {
	return fmt.Sprintf("[GET /dhcps/{dhcpId}][%d] getDhcpOptionsOK  %+v", 200, o.Payload)
}

func (o *GetDhcpOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.DhcpOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpOptionsUnauthorized creates a GetDhcpOptionsUnauthorized with default headers values
func NewGetDhcpOptionsUnauthorized() *GetDhcpOptionsUnauthorized {
	return &GetDhcpOptionsUnauthorized{}
}

/*GetDhcpOptionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDhcpOptionsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetDhcpOptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dhcps/{dhcpId}][%d] getDhcpOptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDhcpOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpOptionsNotFound creates a GetDhcpOptionsNotFound with default headers values
func NewGetDhcpOptionsNotFound() *GetDhcpOptionsNotFound {
	return &GetDhcpOptionsNotFound{}
}

/*GetDhcpOptionsNotFound handles this case with default header values.

Not Found
*/
type GetDhcpOptionsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetDhcpOptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /dhcps/{dhcpId}][%d] getDhcpOptionsNotFound  %+v", 404, o.Payload)
}

func (o *GetDhcpOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpOptionsInternalServerError creates a GetDhcpOptionsInternalServerError with default headers values
func NewGetDhcpOptionsInternalServerError() *GetDhcpOptionsInternalServerError {
	return &GetDhcpOptionsInternalServerError{}
}

/*GetDhcpOptionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDhcpOptionsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetDhcpOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dhcps/{dhcpId}][%d] getDhcpOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDhcpOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpOptionsDefault creates a GetDhcpOptionsDefault with default headers values
func NewGetDhcpOptionsDefault(code int) *GetDhcpOptionsDefault {
	return &GetDhcpOptionsDefault{
		_statusCode: code,
	}
}

/*GetDhcpOptionsDefault handles this case with default header values.

An error has occurred.
*/
type GetDhcpOptionsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get dhcp options default response
func (o *GetDhcpOptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetDhcpOptionsDefault) Error() string {
	return fmt.Sprintf("[GET /dhcps/{dhcpId}][%d] GetDhcpOptions default  %+v", o._statusCode, o.Payload)
}

func (o *GetDhcpOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}