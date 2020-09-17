// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// GetDialogOKCode is the HTTP code returned for type GetDialogOK
const GetDialogOKCode int = 200

/*GetDialogOK return dialogs

swagger:response getDialogOK
*/
type GetDialogOK struct {

	/*
	  In: Body
	*/
	Payload *GetDialogOKBody `json:"body,omitempty"`
}

// NewGetDialogOK creates GetDialogOK with default headers values
func NewGetDialogOK() *GetDialogOK {

	return &GetDialogOK{}
}

// WithPayload adds the payload to the get dialog o k response
func (o *GetDialogOK) WithPayload(payload *GetDialogOKBody) *GetDialogOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog o k response
func (o *GetDialogOK) SetPayload(payload *GetDialogOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDialogBadRequestCode is the HTTP code returned for type GetDialogBadRequest
const GetDialogBadRequestCode int = 400

/*GetDialogBadRequest request error

swagger:response getDialogBadRequest
*/
type GetDialogBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogBadRequest creates GetDialogBadRequest with default headers values
func NewGetDialogBadRequest() *GetDialogBadRequest {

	return &GetDialogBadRequest{}
}

// WithPayload adds the payload to the get dialog bad request response
func (o *GetDialogBadRequest) WithPayload(payload models.Error) *GetDialogBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog bad request response
func (o *GetDialogBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDialogInternalServerErrorCode is the HTTP code returned for type GetDialogInternalServerError
const GetDialogInternalServerErrorCode int = 500

/*GetDialogInternalServerError internal serever error

swagger:response getDialogInternalServerError
*/
type GetDialogInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogInternalServerError creates GetDialogInternalServerError with default headers values
func NewGetDialogInternalServerError() *GetDialogInternalServerError {

	return &GetDialogInternalServerError{}
}

// WithPayload adds the payload to the get dialog internal server error response
func (o *GetDialogInternalServerError) WithPayload(payload models.Error) *GetDialogInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog internal server error response
func (o *GetDialogInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetDialogDefault generic error response

swagger:response getDialogDefault
*/
type GetDialogDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogDefault creates GetDialogDefault with default headers values
func NewGetDialogDefault(code int) *GetDialogDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDialogDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dialog default response
func (o *GetDialogDefault) WithStatusCode(code int) *GetDialogDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dialog default response
func (o *GetDialogDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get dialog default response
func (o *GetDialogDefault) WithPayload(payload models.Error) *GetDialogDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog default response
func (o *GetDialogDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
