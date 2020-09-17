// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostDialogHandlerFunc turns a function with the right signature into a post dialog handler
type PostDialogHandlerFunc func(PostDialogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostDialogHandlerFunc) Handle(params PostDialogParams) middleware.Responder {
	return fn(params)
}

// PostDialogHandler interface for that can handle valid post dialog params
type PostDialogHandler interface {
	Handle(PostDialogParams) middleware.Responder
}

// NewPostDialog creates a new http.Handler for the post dialog operation
func NewPostDialog(ctx *middleware.Context, handler PostDialogHandler) *PostDialog {
	return &PostDialog{Context: ctx, Handler: handler}
}

/*PostDialog swagger:route POST /dialog postDialog

PostDialog post dialog API

*/
type PostDialog struct {
	Context *middleware.Context
	Handler PostDialogHandler
}

func (o *PostDialog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostDialogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostDialogBody post dialog body
//
// swagger:model PostDialogBody
type PostDialogBody struct {

	// author
	Author string `json:"author,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// style
	Style string `json:"style,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (o *PostDialogBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// author
		Author string `json:"author,omitempty"`

		// comment
		Comment string `json:"comment,omitempty"`

		// content
		Content string `json:"content,omitempty"`

		// link
		Link string `json:"link,omitempty"`

		// style
		Style string `json:"style,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// user id
		UserID int64 `json:"user_id,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	o.Author = props.Author
	o.Comment = props.Comment
	o.Content = props.Content
	o.Link = props.Link
	o.Style = props.Style
	o.Title = props.Title
	o.UserID = props.UserID
	return nil
}

// Validate validates this post dialog body
func (o *PostDialogBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDialogBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDialogBody) UnmarshalBinary(b []byte) error {
	var res PostDialogBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
