package config

import (
	"context"
	"database/sql"

	users "github.com/adavarski/Go-gRPC-app-opentelemetry-example/grpc-server/service"
	"github.com/adavarski/Go-gRPC-app-opentelemetry-example/http-server/telemetry"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/blob"
)

type AppConfig struct {
	PackageBucket *blob.Bucket
	Db            *sql.DB
	UsersSvc      users.UsersClient

	// telemetry
	Logger   zerolog.Logger
	Metrics  telemetry.MetricReporter
	Trace    trace.Tracer
	TraceCtx context.Context
	Span     trace.Span
	SpanCtx  context.Context
}
