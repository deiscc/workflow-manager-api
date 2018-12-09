package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*PingOK server ping success

swagger:response pingOK
*/
type PingOK struct {
}

// NewPingOK creates PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

// WriteResponse to the client
func (o *PingOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
}

/*PingDefault unexpected error

swagger:response pingDefault
*/
type PingDefault struct {
	_statusCode int
}

// NewPingDefault creates PingDefault with default headers values
func NewPingDefault(code int) *PingDefault {
	if code <= 0 {
		code = 500
	}

	return &PingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the ping default response
func (o *PingDefault) WithStatusCode(code int) *PingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the ping default response
func (o *PingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PingDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
}