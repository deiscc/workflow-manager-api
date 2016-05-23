package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
	spec "github.com/go-swagger/go-swagger/spec"
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewWorkflowManagerAPI creates a new WorkflowManager instance
func NewWorkflowManagerAPI(spec *spec.Document) *WorkflowManagerAPI {
	o := &WorkflowManagerAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*WorkflowManagerAPI the workflow manager API */
type WorkflowManagerAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer httpkit.Producer

	// CreateClusterDetailsHandler sets the operation handler for the create cluster details operation
	CreateClusterDetailsHandler CreateClusterDetailsHandler
	// GetClusterByIDHandler sets the operation handler for the get cluster by id operation
	GetClusterByIDHandler GetClusterByIDHandler
	// GetClustersByAgeHandler sets the operation handler for the get clusters by age operation
	GetClustersByAgeHandler GetClustersByAgeHandler
	// GetClustersCountHandler sets the operation handler for the get clusters count operation
	GetClustersCountHandler GetClustersCountHandler
	// GetComponentByNameHandler sets the operation handler for the get component by name operation
	GetComponentByNameHandler GetComponentByNameHandler
	// GetComponentByReleaseHandler sets the operation handler for the get component by release operation
	GetComponentByReleaseHandler GetComponentByReleaseHandler
	// GetComponentsByLatestReleaseHandler sets the operation handler for the get components by latest release operation
	GetComponentsByLatestReleaseHandler GetComponentsByLatestReleaseHandler
	// PublishComponentReleaseHandler sets the operation handler for the publish component release operation
	PublishComponentReleaseHandler PublishComponentReleaseHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
}

// SetDefaultProduces sets the default produces media type
func (o *WorkflowManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WorkflowManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *WorkflowManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WorkflowManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WorkflowManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WorkflowManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WorkflowManagerAPI
func (o *WorkflowManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CreateClusterDetailsHandler == nil {
		unregistered = append(unregistered, "CreateClusterDetailsHandler")
	}

	if o.GetClusterByIDHandler == nil {
		unregistered = append(unregistered, "GetClusterByIDHandler")
	}

	if o.GetClustersByAgeHandler == nil {
		unregistered = append(unregistered, "GetClustersByAgeHandler")
	}

	if o.GetClustersCountHandler == nil {
		unregistered = append(unregistered, "GetClustersCountHandler")
	}

	if o.GetComponentByNameHandler == nil {
		unregistered = append(unregistered, "GetComponentByNameHandler")
	}

	if o.GetComponentByReleaseHandler == nil {
		unregistered = append(unregistered, "GetComponentByReleaseHandler")
	}

	if o.GetComponentsByLatestReleaseHandler == nil {
		unregistered = append(unregistered, "GetComponentsByLatestReleaseHandler")
	}

	if o.PublishComponentReleaseHandler == nil {
		unregistered = append(unregistered, "PublishComponentReleaseHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WorkflowManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WorkflowManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *WorkflowManagerAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *WorkflowManagerAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WorkflowManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *WorkflowManagerAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = NewCreateClusterDetails(o.context, o.CreateClusterDetailsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{id}"] = NewGetClusterByID(o.context, o.GetClusterByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/age"] = NewGetClustersByAge(o.context, o.GetClustersByAgeHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/count"] = NewGetClustersCount(o.context, o.GetClustersCountHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/versions/{train}/{component}"] = NewGetComponentByName(o.context, o.GetComponentByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/versions/{train}/{component}/{release}"] = NewGetComponentByRelease(o.context, o.GetComponentByReleaseHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/versions/latest"] = NewGetComponentsByLatestRelease(o.context, o.GetComponentsByLatestReleaseHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/versions/{train}/{component}/{release}"] = NewPublishComponentRelease(o.context, o.PublishComponentReleaseHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WorkflowManagerAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
