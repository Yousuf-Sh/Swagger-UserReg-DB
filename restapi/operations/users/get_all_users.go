// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllUsersHandlerFunc turns a function with the right signature into a get all users handler
type GetAllUsersHandlerFunc func(GetAllUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllUsersHandlerFunc) Handle(params GetAllUsersParams) middleware.Responder {
	return fn(params)
}

// GetAllUsersHandler interface for that can handle valid get all users params
type GetAllUsersHandler interface {
	Handle(GetAllUsersParams) middleware.Responder
}

// NewGetAllUsers creates a new http.Handler for the get all users operation
func NewGetAllUsers(ctx *middleware.Context, handler GetAllUsersHandler) *GetAllUsers {
	return &GetAllUsers{Context: ctx, Handler: handler}
}

/*
	GetAllUsers swagger:route GET /users users getAllUsers

Get all users
*/
type GetAllUsers struct {
	Context *middleware.Context
	Handler GetAllUsersHandler
}

func (o *GetAllUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllUsersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
