package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
)

/*GetPersistentClustersOK clusters details response

swagger:response getPersistentClustersOK
*/
type GetPersistentClustersOK struct {

	// In: body
	Payload *models.ClustersCount `json:"body,omitempty"`
}

// NewGetPersistentClustersOK creates GetPersistentClustersOK with default headers values
func NewGetPersistentClustersOK() *GetPersistentClustersOK {
	return &GetPersistentClustersOK{}
}

// WithPayload adds the payload to the get persistent clusters o k response
func (o *GetPersistentClustersOK) WithPayload(payload *models.ClustersCount) *GetPersistentClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get persistent clusters o k response
func (o *GetPersistentClustersOK) SetPayload(payload *models.ClustersCount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersistentClustersOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPersistentClustersDefault unexpected error

swagger:response getPersistentClustersDefault
*/
type GetPersistentClustersDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPersistentClustersDefault creates GetPersistentClustersDefault with default headers values
func NewGetPersistentClustersDefault(code int) *GetPersistentClustersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPersistentClustersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get persistent clusters default response
func (o *GetPersistentClustersDefault) WithStatusCode(code int) *GetPersistentClustersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get persistent clusters default response
func (o *GetPersistentClustersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get persistent clusters default response
func (o *GetPersistentClustersDefault) WithPayload(payload *models.Error) *GetPersistentClustersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get persistent clusters default response
func (o *GetPersistentClustersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersistentClustersDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
