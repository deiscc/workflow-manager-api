package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
)

/*CreateClusterDetailsForV2OK clusters details response

swagger:response createClusterDetailsForV2OK
*/
type CreateClusterDetailsForV2OK struct {

	// In: body
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewCreateClusterDetailsForV2OK creates CreateClusterDetailsForV2OK with default headers values
func NewCreateClusterDetailsForV2OK() *CreateClusterDetailsForV2OK {
	return &CreateClusterDetailsForV2OK{}
}

// WithPayload adds the payload to the create cluster details for v2 o k response
func (o *CreateClusterDetailsForV2OK) WithPayload(payload *models.Cluster) *CreateClusterDetailsForV2OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster details for v2 o k response
func (o *CreateClusterDetailsForV2OK) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterDetailsForV2OK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateClusterDetailsForV2Default unexpected error

swagger:response createClusterDetailsForV2Default
*/
type CreateClusterDetailsForV2Default struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateClusterDetailsForV2Default creates CreateClusterDetailsForV2Default with default headers values
func NewCreateClusterDetailsForV2Default(code int) *CreateClusterDetailsForV2Default {
	if code <= 0 {
		code = 500
	}

	return &CreateClusterDetailsForV2Default{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create cluster details for v2 default response
func (o *CreateClusterDetailsForV2Default) WithStatusCode(code int) *CreateClusterDetailsForV2Default {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create cluster details for v2 default response
func (o *CreateClusterDetailsForV2Default) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create cluster details for v2 default response
func (o *CreateClusterDetailsForV2Default) WithPayload(payload *models.Error) *CreateClusterDetailsForV2Default {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster details for v2 default response
func (o *CreateClusterDetailsForV2Default) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterDetailsForV2Default) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}