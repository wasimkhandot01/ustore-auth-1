// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProfileHandlerFunc turns a function with the right signature into a profile handler
type ProfileHandlerFunc func(ProfileParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ProfileHandlerFunc) Handle(params ProfileParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ProfileHandler interface for that can handle valid profile params
type ProfileHandler interface {
	Handle(ProfileParams, interface{}) middleware.Responder
}

// NewProfile creates a new http.Handler for the profile operation
func NewProfile(ctx *middleware.Context, handler ProfileHandler) *Profile {
	return &Profile{Context: ctx, Handler: handler}
}

/* Profile swagger:route GET /user/profile User profile

To show user details

*/
type Profile struct {
	Context *middleware.Context
	Handler ProfileHandler
}

func (o *Profile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewProfileParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
