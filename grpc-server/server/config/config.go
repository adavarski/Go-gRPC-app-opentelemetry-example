package config

import (
	"github.com/adavarski/Go-gRPC-app-opentelemetry-example/grpc-server/server/telemetry"
	"github.com/rs/zerolog"
)

type AppConfig struct {
	Logger  zerolog.Logger
	Metrics telemetry.MetricReporter
}
