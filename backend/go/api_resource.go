/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package shellbackend

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ResourceAPIController binds http requests to an api service and writes the service results to the http response
type ResourceAPIController struct {
	service      ResourceAPIServicer
	errorHandler ErrorHandler
}

// ResourceAPIOption for how the controller is set up.
type ResourceAPIOption func(*ResourceAPIController)

// WithResourceAPIErrorHandler inject ErrorHandler into controller
func WithResourceAPIErrorHandler(h ErrorHandler) ResourceAPIOption {
	return func(c *ResourceAPIController) {
		c.errorHandler = h
	}
}

// NewResourceAPIController creates a default api controller
func NewResourceAPIController(s ResourceAPIServicer, opts ...ResourceAPIOption) Router {
	controller := &ResourceAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ResourceAPIController
func (c *ResourceAPIController) Routes() Routes {
	return Routes{
		"GetResourceById": Route{
			strings.ToUpper("Get"),
			"/api/v3/resource/{resourceId}",
			c.GetResourceById,
		},
		"GetResources": Route{
			strings.ToUpper("Get"),
			"/api/v3/resource/",
			c.GetResources,
		},
	}
}

// GetResourceById - Find resource by ID
func (c *ResourceAPIController) GetResourceById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resourceIdParam, err := parseNumericParameter[int64](
		params["resourceId"],
		WithRequire[int64](parseInt64),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetResourceById(r.Context(), resourceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetResources - Returns a list of resources
func (c *ResourceAPIController) GetResources(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetResources(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
