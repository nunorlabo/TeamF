// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/gen/models"
)

// GetDialogIDCommentOKCode is the HTTP code returned for type GetDialogIDCommentOK
const GetDialogIDCommentOKCode int = 200

/*GetDialogIDCommentOK 取得成功

swagger:response getDialogIdCommentOK
*/
type GetDialogIDCommentOK struct {

	/*
	  In: Body
	*/
	Payload *GetDialogIDCommentOKBody `json:"body,omitempty"`
}

// NewGetDialogIDCommentOK creates GetDialogIDCommentOK with default headers values
func NewGetDialogIDCommentOK() *GetDialogIDCommentOK {

	return &GetDialogIDCommentOK{}
}

// WithPayload adds the payload to the get dialog Id comment o k response
func (o *GetDialogIDCommentOK) WithPayload(payload *GetDialogIDCommentOKBody) *GetDialogIDCommentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog Id comment o k response
func (o *GetDialogIDCommentOK) SetPayload(payload *GetDialogIDCommentOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogIDCommentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDialogIDCommentBadRequestCode is the HTTP code returned for type GetDialogIDCommentBadRequest
const GetDialogIDCommentBadRequestCode int = 400

/*GetDialogIDCommentBadRequest request error

swagger:response getDialogIdCommentBadRequest
*/
type GetDialogIDCommentBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogIDCommentBadRequest creates GetDialogIDCommentBadRequest with default headers values
func NewGetDialogIDCommentBadRequest() *GetDialogIDCommentBadRequest {

	return &GetDialogIDCommentBadRequest{}
}

// WithPayload adds the payload to the get dialog Id comment bad request response
func (o *GetDialogIDCommentBadRequest) WithPayload(payload models.Error) *GetDialogIDCommentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog Id comment bad request response
func (o *GetDialogIDCommentBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogIDCommentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDialogIDCommentInternalServerErrorCode is the HTTP code returned for type GetDialogIDCommentInternalServerError
const GetDialogIDCommentInternalServerErrorCode int = 500

/*GetDialogIDCommentInternalServerError internal serever error

swagger:response getDialogIdCommentInternalServerError
*/
type GetDialogIDCommentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogIDCommentInternalServerError creates GetDialogIDCommentInternalServerError with default headers values
func NewGetDialogIDCommentInternalServerError() *GetDialogIDCommentInternalServerError {

	return &GetDialogIDCommentInternalServerError{}
}

// WithPayload adds the payload to the get dialog Id comment internal server error response
func (o *GetDialogIDCommentInternalServerError) WithPayload(payload models.Error) *GetDialogIDCommentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog Id comment internal server error response
func (o *GetDialogIDCommentInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogIDCommentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetDialogIDCommentDefault generic error response

swagger:response getDialogIdCommentDefault
*/
type GetDialogIDCommentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetDialogIDCommentDefault creates GetDialogIDCommentDefault with default headers values
func NewGetDialogIDCommentDefault(code int) *GetDialogIDCommentDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDialogIDCommentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dialog ID comment default response
func (o *GetDialogIDCommentDefault) WithStatusCode(code int) *GetDialogIDCommentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dialog ID comment default response
func (o *GetDialogIDCommentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get dialog ID comment default response
func (o *GetDialogIDCommentDefault) WithPayload(payload models.Error) *GetDialogIDCommentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dialog ID comment default response
func (o *GetDialogIDCommentDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDialogIDCommentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
