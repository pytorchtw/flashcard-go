// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetGreetingOKCode is the HTTP code returned for type GetGreetingOK
const GetGreetingOKCode int = 200

/*GetGreetingOK returns a greeting

swagger:response getGreetingOK
*/
type GetGreetingOK struct {
}

// NewGetGreetingOK creates GetGreetingOK with default headers values
func NewGetGreetingOK() *GetGreetingOK {

	return &GetGreetingOK{}
}

// WriteResponse to the client
func (o *GetGreetingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
