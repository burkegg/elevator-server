package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/onbeep/elevator-server/go-ele/restapi/operations"
	"github.com/onbeep/elevator-server/go-ele/vator"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../../elevator.yml

func configureFlags(api *operations.LiftyAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LiftyAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Declare the object for the Elevator
	v, err := vator.NewVator([]string{"B1", "F1", "F2", "F3"}, 3)
	if err != nil {
		panic(err)
	}

	api.CurrentFloorHandler = operations.CurrentFloorHandlerFunc(func(params operations.CurrentFloorParams) middleware.Responder {
		f := v.Current(params.CarID)
		body := operations.CurrentFloorOKBody{ID: &f.ID, Name: &f.Name}
		return operations.NewCurrentFloorOK().WithPayload(body)
	})
	api.FloorCountHandler = operations.FloorCountHandlerFunc(func(params operations.FloorCountParams) middleware.Responder {
		// return middleware.NotImplemented("operation .FloorCount has not yet been implemented")
		length := int32(len(v.Floors()))
		return operations.NewFloorCountOK().WithPayload(operations.FloorCountOKBody{
			Count: &length,
		})
	})
	api.InventoryHandler = operations.InventoryHandlerFunc(func(params operations.InventoryParams) middleware.Responder {
		if params.Pwd != "p4ssw3rd" {
			return operations.NewInventoryUnauthorized()
		}

		body := []*operations.InventoryOKBodyItems0{}
		for _, floor := range v.Floors() {
			fid := floor.ID
			fname := floor.Name
			body = append(body, &operations.InventoryOKBodyItems0{
				ID:   fid,
				Name: fname,
			})
		}
		for _, car := range v.Cars() {
			cid := car.ID
			cname := car.Name
			body = append(body, &operations.InventoryOKBodyItems0{
				ID:   cid,
				Name: cname,
			})
		}

		return operations.NewInventoryOK().WithPayload(body)
	})
	api.WelcomeHandler = operations.WelcomeHandlerFunc(func(params operations.WelcomeParams) middleware.Responder {
		w := "Welcome to the Elevator Server"
		return operations.NewWelcomeOK().WithPayload(operations.WelcomeOKBody{Msg: &w})
	})

	api.ServerShutdown = func() {
		fmt.Printf("Shutting Down!\n")
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
