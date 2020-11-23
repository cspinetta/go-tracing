package base

import "go.opentelemetry.io/otel"

var GlobalAppTracer = otel.Tracer("go-tracing")
