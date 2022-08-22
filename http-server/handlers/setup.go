package handlers

import (
	"net/http"

	"github.com/adavarski/Go-gRPC-app-opentelemetry-example/http-server/config"
	"github.com/adavarski/Go-gRPC-app-opentelemetry-example/http-server/types"
)

func SetupHandlers(mux *http.ServeMux, config *config.AppConfig) {
	mux.Handle(
		"/api/packages",
		&types.App{Config: config, Handler: packageHandler},
	)
}
