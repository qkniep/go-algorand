// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":           true,
		"fee":              true,
		"key-dilution":     true,
		"round-last-valid": true,
		"no-wait":          true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
