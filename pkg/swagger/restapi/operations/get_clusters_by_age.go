package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetClustersByAgeHandlerFunc turns a function with the right signature into a get clusters by age handler
type GetClustersByAgeHandlerFunc func(GetClustersByAgeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClustersByAgeHandlerFunc) Handle(params GetClustersByAgeParams) middleware.Responder {
	return fn(params)
}

// GetClustersByAgeHandler interface for that can handle valid get clusters by age params
type GetClustersByAgeHandler interface {
	Handle(GetClustersByAgeParams) middleware.Responder
}

// NewGetClustersByAge creates a new http.Handler for the get clusters by age operation
func NewGetClustersByAge(ctx *middleware.Context, handler GetClustersByAgeHandler) *GetClustersByAge {
	return &GetClustersByAge{Context: ctx, Handler: handler}
}

/*GetClustersByAge swagger:route GET /v3/clusters/age getClustersByAge

list clusters

*/
type GetClustersByAge struct {
	Context *middleware.Context
	Handler GetClustersByAgeHandler
}

func (o *GetClustersByAge) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetClustersByAgeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetClustersByAgeOKBodyBody get clusters by age o k body body

swagger:model GetClustersByAgeOKBodyBody
*/
type GetClustersByAgeOKBodyBody struct {

	/* data

	Required: true
	*/
	Data []*models.Cluster `json:"data"`
}

// Validate validates this get clusters by age o k body body
func (o *GetClustersByAgeOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetClustersByAgeOKBodyBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getClustersByAgeOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if err := o.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
