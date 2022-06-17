// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api/client/models"
)

// ResonateUserAddUserReader is a Reader for the ResonateUserAddUser structure.
type ResonateUserAddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserAddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserAddUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserAddUserOK creates a ResonateUserAddUserOK with default headers values
func NewResonateUserAddUserOK() *ResonateUserAddUserOK {
	return &ResonateUserAddUserOK{}
}

/* ResonateUserAddUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserAddUserOK struct {
	Payload *models.UserUserRequest
}

func (o *ResonateUserAddUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/users][%d] resonateUserAddUserOK  %+v", 200, o.Payload)
}
func (o *ResonateUserAddUserOK) GetPayload() *models.UserUserRequest {
	return o.Payload
}

func (o *ResonateUserAddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUserRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserAddUserDefault creates a ResonateUserAddUserDefault with default headers values
func NewResonateUserAddUserDefault(code int) *ResonateUserAddUserDefault {
	return &ResonateUserAddUserDefault{
		_statusCode: code,
	}
}

/* ResonateUserAddUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserAddUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user add user default response
func (o *ResonateUserAddUserDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserAddUserDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/users][%d] ResonateUser_AddUser default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserAddUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserAddUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
