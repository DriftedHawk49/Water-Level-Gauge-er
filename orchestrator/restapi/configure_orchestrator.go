// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Water-Level-Gauge-er/orchestrator/restapi/operations"
	"Water-Level-Gauge-er/orchestrator/restapi/operations/configuration"
	"Water-Level-Gauge-er/orchestrator/restapi/operations/water_level"
	"Water-Level-Gauge-er/orchestrator/restapi/operations/water_level_anomaly"
)

//go:generate swagger generate server --target ../../orchestrator --name Orchestrator --spec ../api/swagger_spec.yaml --principal any

func configureFlags(api *operations.OrchestratorAPI) {
	// api.CommandLineOptionsGroups = []cmdutils.CommandLineOptionsGroup{ ... }
	_ = api
}

func configureAPI(api *operations.OrchestratorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...any)
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ConfigurationGetConfigurationHandler == nil {
		api.ConfigurationGetConfigurationHandler = configuration.GetConfigurationHandlerFunc(func(params configuration.GetConfigurationParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation configuration.GetConfiguration has not yet been implemented")
		})
	}
	if api.WaterLevelAnomalyGetWaterAnomalyHandler == nil {
		api.WaterLevelAnomalyGetWaterAnomalyHandler = water_level_anomaly.GetWaterAnomalyHandlerFunc(func(params water_level_anomaly.GetWaterAnomalyParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation water_level_anomaly.GetWaterAnomaly has not yet been implemented")
		})
	}
	if api.WaterLevelAnomalyGetWaterAnomalyHistoryHandler == nil {
		api.WaterLevelAnomalyGetWaterAnomalyHistoryHandler = water_level_anomaly.GetWaterAnomalyHistoryHandlerFunc(func(params water_level_anomaly.GetWaterAnomalyHistoryParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation water_level_anomaly.GetWaterAnomalyHistory has not yet been implemented")
		})
	}
	if api.WaterLevelGetWaterLevelHistoryHandler == nil {
		api.WaterLevelGetWaterLevelHistoryHandler = water_level.GetWaterLevelHistoryHandlerFunc(func(params water_level.GetWaterLevelHistoryParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation water_level.GetWaterLevelHistory has not yet been implemented")
		})
	}
	if api.ConfigurationPostConfigurationHandler == nil {
		api.ConfigurationPostConfigurationHandler = configuration.PostConfigurationHandlerFunc(func(params configuration.PostConfigurationParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation configuration.PostConfiguration has not yet been implemented")
		})
	}
	if api.WaterLevelGetWaterLevelHandler == nil {
		api.WaterLevelGetWaterLevelHandler = water_level.GetWaterLevelHandlerFunc(func(params water_level.GetWaterLevelParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation water_level.GetWaterLevel has not yet been implemented")
		})
	}
	if api.WaterLevelSaveWaterLevelHandler == nil {
		api.WaterLevelSaveWaterLevelHandler = water_level.SaveWaterLevelHandlerFunc(func(params water_level.SaveWaterLevelParams) middleware.Responder {
			_ = params

			return middleware.NotImplemented("operation water_level.SaveWaterLevel has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	_ = tlsConfig
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(server *http.Server, scheme, addr string) {
	_ = server
	_ = scheme
	_ = addr
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
