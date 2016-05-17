package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager-api/pkg/swagger/models"
)

// NewPublishComponentReleaseParams creates a new PublishComponentReleaseParams object
// with the default values initialized.
func NewPublishComponentReleaseParams() PublishComponentReleaseParams {
	var ()
	return PublishComponentReleaseParams{}
}

// PublishComponentReleaseParams contains all the bound params for the publish component release operation
// typically these are obtained from a http.Request
//
// swagger:parameters publishComponentRelease
type PublishComponentReleaseParams struct {
	/*
	  In: body
	*/
	Body *models.ComponentVersion
	/*A component is a single deis component, e.g., deis-router
	  Required: true
	  In: path
	*/
	Component string
	/*The release version of the deis component, eg., 2.0.0-beta2
	  Required: true
	  In: path
	*/
	Release string
	/*A train is a release cadence type, e.g., "beta" or "stable"
	  Required: true
	  In: path
	*/
	Train string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PublishComponentReleaseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	defer r.Body.Close()
	var body models.ComponentVersion
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.Body = &body
		}
	}

	rComponent, rhkComponent, _ := route.Params.GetOK("component")
	if err := o.bindComponent(rComponent, rhkComponent, route.Formats); err != nil {
		res = append(res, err)
	}

	rRelease, rhkRelease, _ := route.Params.GetOK("release")
	if err := o.bindRelease(rRelease, rhkRelease, route.Formats); err != nil {
		res = append(res, err)
	}

	rTrain, rhkTrain, _ := route.Params.GetOK("train")
	if err := o.bindTrain(rTrain, rhkTrain, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PublishComponentReleaseParams) bindComponent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Component = raw

	return nil
}

func (o *PublishComponentReleaseParams) bindRelease(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Release = raw

	return nil
}

func (o *PublishComponentReleaseParams) bindTrain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Train = raw

	return nil
}
