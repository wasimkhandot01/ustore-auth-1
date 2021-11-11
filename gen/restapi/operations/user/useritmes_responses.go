// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wasimkhan042/ustore-auth/gen/models"
)

// UseritmesOKCode is the HTTP code returned for type UseritmesOK
const UseritmesOKCode int = 200

/*UseritmesOK Success response when items are shown

swagger:response useritmesOK
*/
type UseritmesOK struct {

	/*
	  In: Body
	*/
	Payload models.UserProducts `json:"body,omitempty"`
}

// NewUseritmesOK creates UseritmesOK with default headers values
func NewUseritmesOK() *UseritmesOK {

	return &UseritmesOK{}
}

// WithPayload adds the payload to the useritmes o k response
func (o *UseritmesOK) WithPayload(payload models.UserProducts) *UseritmesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the useritmes o k response
func (o *UseritmesOK) SetPayload(payload models.UserProducts) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UseritmesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.UserProducts{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UseritmesBadRequestCode is the HTTP code returned for type UseritmesBadRequest
const UseritmesBadRequestCode int = 400

/*UseritmesBadRequest Bad Request

swagger:response useritmesBadRequest
*/
type UseritmesBadRequest struct {
}

// NewUseritmesBadRequest creates UseritmesBadRequest with default headers values
func NewUseritmesBadRequest() *UseritmesBadRequest {

	return &UseritmesBadRequest{}
}

// WriteResponse to the client
func (o *UseritmesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UseritmesNotFoundCode is the HTTP code returned for type UseritmesNotFound
const UseritmesNotFoundCode int = 404

/*UseritmesNotFound items not found

swagger:response useritmesNotFound
*/
type UseritmesNotFound struct {
}

// NewUseritmesNotFound creates UseritmesNotFound with default headers values
func NewUseritmesNotFound() *UseritmesNotFound {

	return &UseritmesNotFound{}
}

// WriteResponse to the client
func (o *UseritmesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UseritmesInternalServerErrorCode is the HTTP code returned for type UseritmesInternalServerError
const UseritmesInternalServerErrorCode int = 500

/*UseritmesInternalServerError Server error

swagger:response useritmesInternalServerError
*/
type UseritmesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUseritmesInternalServerError creates UseritmesInternalServerError with default headers values
func NewUseritmesInternalServerError() *UseritmesInternalServerError {

	return &UseritmesInternalServerError{}
}

// WithPayload adds the payload to the useritmes internal server error response
func (o *UseritmesInternalServerError) WithPayload(payload string) *UseritmesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the useritmes internal server error response
func (o *UseritmesInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UseritmesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}