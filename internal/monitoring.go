package internal

import "go.opentelemetry.io/otel"

var (
	Tracer = otel.Tracer("skeleton")
	Meter  = otel.Meter("skeleton")
)
