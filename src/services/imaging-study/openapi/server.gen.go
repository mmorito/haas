// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create new DicomStore
	// (POST /v1/organizations/{organizationID}/dicomStore)
	CreateDicomStore(ctx echo.Context, organizationID string) error
	// Store Instances (STOW-RS)
	// (POST /v1/organizations/{organizationID}/studies)
	StoreInstances(ctx echo.Context, organizationID string) error
	// Retrieve Study (WADO-RS)
	// (GET /v1/organizations/{organizationID}/studies/{studyInstanceUID})
	RetrieveStudy(ctx echo.Context, organizationID string, studyInstanceUID string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateDicomStore converts echo context to params.
func (w *ServerInterfaceWrapper) CreateDicomStore(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "organizationID" -------------
	var organizationID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, ctx.Param("organizationID"), &organizationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter organizationID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateDicomStore(ctx, organizationID)
	return err
}

// StoreInstances converts echo context to params.
func (w *ServerInterfaceWrapper) StoreInstances(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "organizationID" -------------
	var organizationID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, ctx.Param("organizationID"), &organizationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter organizationID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StoreInstances(ctx, organizationID)
	return err
}

// RetrieveStudy converts echo context to params.
func (w *ServerInterfaceWrapper) RetrieveStudy(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "organizationID" -------------
	var organizationID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "organizationID", runtime.ParamLocationPath, ctx.Param("organizationID"), &organizationID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter organizationID: %s", err))
	}

	// ------------- Path parameter "studyInstanceUID" -------------
	var studyInstanceUID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "studyInstanceUID", runtime.ParamLocationPath, ctx.Param("studyInstanceUID"), &studyInstanceUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter studyInstanceUID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RetrieveStudy(ctx, organizationID, studyInstanceUID)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v1/organizations/:organizationID/dicomStore", wrapper.CreateDicomStore)
	router.POST(baseURL+"/v1/organizations/:organizationID/studies", wrapper.StoreInstances)
	router.GET(baseURL+"/v1/organizations/:organizationID/studies/:studyInstanceUID", wrapper.RetrieveStudy)

}
