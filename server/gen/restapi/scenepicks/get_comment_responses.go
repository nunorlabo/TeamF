// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// GetCommentOKCode is the HTTP code returned for type GetCommentOK
const GetCommentOKCode int = 200

/*GetCommentOK 取得成功

swagger:response getCommentOK
*/
type GetCommentOK struct {

	/*
	  In: Body
	*/
	Payload *GetCommentOKBody `json:"body,omitempty"`
}

// NewGetCommentOK creates GetCommentOK with default headers values
func NewGetCommentOK() *GetCommentOK {

	return &GetCommentOK{}
}

// WithPayload adds the payload to the get comment o k response
func (o *GetCommentOK) WithPayload(payload *GetCommentOKBody) *GetCommentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment o k response
func (o *GetCommentOK) SetPayload(payload *GetCommentOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCommentBadRequestCode is the HTTP code returned for type GetCommentBadRequest
const GetCommentBadRequestCode int = 400

/*GetCommentBadRequest request error

swagger:response getCommentBadRequest
*/
type GetCommentBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentBadRequest creates GetCommentBadRequest with default headers values
func NewGetCommentBadRequest() *GetCommentBadRequest {

	return &GetCommentBadRequest{}
}

// WithPayload adds the payload to the get comment bad request response
func (o *GetCommentBadRequest) WithPayload(payload models.Error) *GetCommentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment bad request response
func (o *GetCommentBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCommentInternalServerErrorCode is the HTTP code returned for type GetCommentInternalServerError
const GetCommentInternalServerErrorCode int = 500

/*GetCommentInternalServerError internal serever error

swagger:response getCommentInternalServerError
*/
type GetCommentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentInternalServerError creates GetCommentInternalServerError with default headers values
func NewGetCommentInternalServerError() *GetCommentInternalServerError {

	return &GetCommentInternalServerError{}
}

// WithPayload adds the payload to the get comment internal server error response
func (o *GetCommentInternalServerError) WithPayload(payload models.Error) *GetCommentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment internal server error response
func (o *GetCommentInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetCommentDefault generic error response

swagger:response getCommentDefault
*/
type GetCommentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentDefault creates GetCommentDefault with default headers values
func NewGetCommentDefault(code int) *GetCommentDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCommentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get comment default response
func (o *GetCommentDefault) WithStatusCode(code int) *GetCommentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get comment default response
func (o *GetCommentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get comment default response
func (o *GetCommentDefault) WithPayload(payload models.Error) *GetCommentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment default response
func (o *GetCommentDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
