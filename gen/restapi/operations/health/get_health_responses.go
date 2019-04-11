// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/daesu/payments/gen/models"
)

// GetHealthOKCode is the HTTP code returned for type GetHealthOK
const GetHealthOKCode int = 200

/*GetHealthOK Success

swagger:response getHealthOK
*/
type GetHealthOK struct {

	/*
	  In: Body
	*/
	Payload *models.Health `json:"body,omitempty"`
}

// NewGetHealthOK creates GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {

	return &GetHealthOK{}
}

// WithPayload adds the payload to the get health o k response
func (o *GetHealthOK) WithPayload(payload *models.Health) *GetHealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health o k response
func (o *GetHealthOK) SetPayload(payload *models.Health) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthBadRequestCode is the HTTP code returned for type GetHealthBadRequest
const GetHealthBadRequestCode int = 400

/*GetHealthBadRequest Invalid request

swagger:response getHealthBadRequest
*/
type GetHealthBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetHealthBadRequest creates GetHealthBadRequest with default headers values
func NewGetHealthBadRequest() *GetHealthBadRequest {

	return &GetHealthBadRequest{}
}

// WithPayload adds the payload to the get health bad request response
func (o *GetHealthBadRequest) WithPayload(payload *models.ErrorResponse) *GetHealthBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health bad request response
func (o *GetHealthBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}