// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// GetCommentByIDOKCode is the HTTP code returned for type GetCommentByIDOK
const GetCommentByIDOKCode int = 200

/*GetCommentByIDOK 取得成功

swagger:response getCommentByIdOK
*/
type GetCommentByIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetCommentByIDOKBody `json:"body,omitempty"`
}

// NewGetCommentByIDOK creates GetCommentByIDOK with default headers values
func NewGetCommentByIDOK() *GetCommentByIDOK {

	return &GetCommentByIDOK{}
}

// WithPayload adds the payload to the get comment by Id o k response
func (o *GetCommentByIDOK) WithPayload(payload *GetCommentByIDOKBody) *GetCommentByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment by Id o k response
func (o *GetCommentByIDOK) SetPayload(payload *GetCommentByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCommentByIDBadRequestCode is the HTTP code returned for type GetCommentByIDBadRequest
const GetCommentByIDBadRequestCode int = 400

/*GetCommentByIDBadRequest request error

swagger:response getCommentByIdBadRequest
*/
type GetCommentByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentByIDBadRequest creates GetCommentByIDBadRequest with default headers values
func NewGetCommentByIDBadRequest() *GetCommentByIDBadRequest {

	return &GetCommentByIDBadRequest{}
}

// WithPayload adds the payload to the get comment by Id bad request response
func (o *GetCommentByIDBadRequest) WithPayload(payload models.Error) *GetCommentByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment by Id bad request response
func (o *GetCommentByIDBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCommentByIDInternalServerErrorCode is the HTTP code returned for type GetCommentByIDInternalServerError
const GetCommentByIDInternalServerErrorCode int = 500

/*GetCommentByIDInternalServerError internal serever error

swagger:response getCommentByIdInternalServerError
*/
type GetCommentByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentByIDInternalServerError creates GetCommentByIDInternalServerError with default headers values
func NewGetCommentByIDInternalServerError() *GetCommentByIDInternalServerError {

	return &GetCommentByIDInternalServerError{}
}

// WithPayload adds the payload to the get comment by Id internal server error response
func (o *GetCommentByIDInternalServerError) WithPayload(payload models.Error) *GetCommentByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment by Id internal server error response
func (o *GetCommentByIDInternalServerError) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetCommentByIDDefault generic error response

swagger:response getCommentByIdDefault
*/
type GetCommentByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetCommentByIDDefault creates GetCommentByIDDefault with default headers values
func NewGetCommentByIDDefault(code int) *GetCommentByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCommentByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get comment by Id default response
func (o *GetCommentByIDDefault) WithStatusCode(code int) *GetCommentByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get comment by Id default response
func (o *GetCommentByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get comment by Id default response
func (o *GetCommentByIDDefault) WithPayload(payload models.Error) *GetCommentByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get comment by Id default response
func (o *GetCommentByIDDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommentByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
