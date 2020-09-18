// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// PostTagOKCode is the HTTP code returned for type PostTagOK
const PostTagOKCode int = 200

/*PostTagOK 取得成功

swagger:response postTagOK
*/
type PostTagOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostTagOK creates PostTagOK with default headers values
func NewPostTagOK() *PostTagOK {

	return &PostTagOK{}
}

// WithPayload adds the payload to the post tag o k response
func (o *PostTagOK) WithPayload(payload string) *PostTagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tag o k response
func (o *PostTagOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostTagBadRequestCode is the HTTP code returned for type PostTagBadRequest
const PostTagBadRequestCode int = 400

/*PostTagBadRequest request error

swagger:response postTagBadRequest
*/
type PostTagBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostTagBadRequest creates PostTagBadRequest with default headers values
func NewPostTagBadRequest() *PostTagBadRequest {

	return &PostTagBadRequest{}
}

// WithPayload adds the payload to the post tag bad request response
func (o *PostTagBadRequest) WithPayload(payload models.Error) *PostTagBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tag bad request response
func (o *PostTagBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTagBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostTagInternalServerErrorCode is the HTTP code returned for type PostTagInternalServerError
const PostTagInternalServerErrorCode int = 500

/*PostTagInternalServerError internal serever error

swagger:response postTagInternalServerError
*/
type PostTagInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostTagInternalServerError creates PostTagInternalServerError with default headers values
func NewPostTagInternalServerError() *PostTagInternalServerError {

	return &PostTagInternalServerError{}
}

// WithPayload adds the payload to the post tag internal server error response
func (o *PostTagInternalServerError) WithPayload(payload models.Error) *PostTagInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tag internal server error response
func (o *PostTagInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTagInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostTagDefault generic error response

swagger:response postTagDefault
*/
type PostTagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostTagDefault creates PostTagDefault with default headers values
func NewPostTagDefault(code int) *PostTagDefault {
	if code <= 0 {
		code = 500
	}

	return &PostTagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post tag default response
func (o *PostTagDefault) WithStatusCode(code int) *PostTagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post tag default response
func (o *PostTagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post tag default response
func (o *PostTagDefault) WithPayload(payload models.Error) *PostTagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post tag default response
func (o *PostTagDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
