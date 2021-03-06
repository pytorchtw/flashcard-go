// Code generated by go-swagger; DO NOT EDIT.

package deck

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LoadDeckHandlerFunc turns a function with the right signature into a load deck handler
type LoadDeckHandlerFunc func(LoadDeckParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoadDeckHandlerFunc) Handle(params LoadDeckParams) middleware.Responder {
	return fn(params)
}

// LoadDeckHandler interface for that can handle valid load deck params
type LoadDeckHandler interface {
	Handle(LoadDeckParams) middleware.Responder
}

// NewLoadDeck creates a new http.Handler for the load deck operation
func NewLoadDeck(ctx *middleware.Context, handler LoadDeckHandler) *LoadDeck {
	return &LoadDeck{Context: ctx, Handler: handler}
}

/*LoadDeck swagger:route POST /deck deck loadDeck

LoadDeck load deck API

*/
type LoadDeck struct {
	Context *middleware.Context
	Handler LoadDeckHandler
}

func (o *LoadDeck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoadDeckParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
