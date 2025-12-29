package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-logr/stdr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	lemonsKey   = attribute.Key("ex.com/lemons")
	fooKey      = attribute.Key("ex.com/foo")
	barKey      = attribute.Key("ex.com/bar")
	anotherKey  = attribute.Key("ex.com/another")
	serviceName = semconv.ServiceNameKey.String("test-service")
)
var tp *sdktrace.TracerProvider

// SubOperation is an example to demonstrate the use of named tracer.
// It creates a named tracer withis package paths
func SubOperation(ctx context.Context) error {
	//Using global provider. Alternative is to have application provide a getter
	//for its component to get the instance of the providerconst
	tr := otel.Tracer("go.opentelemetry.io/contrib/examples/namedtracer/foo")
	var span trace.Span
	_, span = tr.Start(ctx, "Sub operation...")
	defer span.End()
	span.SetAttributes(lemonsKey.String("five"))
	span.AddEvent("Sub span event")
	return nil
}

// initialize a gRPC connection to be used by both the tracer and meter providers
func initConn() (*grpc.ClientConn, error) {
	//It connects the OpenTelemetry Collector through local gRPC connection
	//You may replace `localhost:4317` with your endpoint
	conn, err := grpc.NewClient("localhost:4317",
		//Note the use of insecure transport here. TLS is recommended in production
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}
	return conn, err
}

// Initializes an OTLP exporter, and configures the corresponding trace provider
func initTracerProvider(ctx context.Context,
	res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	//Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	//Register the trace exporter with a TraerProvider, using a batch
	//span processor to aggregate spans before export
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	//Set global propagator to tracecontext (the default is no-op)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	//Shutdown will flush any remaining spans and shut down the exporter
	return tracerProvider.Shutdown, nil
}

func main() {
	//Set logging level to info to see SDK status messages
	stdr.SetVerbosity(5)
	log.Printf("Waiting for connection...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	conn, err := initConn()
	if err != nil {
		log.Fatal(err)
	}
	res, err := resource.New(ctx,
		resource.WithAttributes(
			//The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	shutdownTracerProvider, err := initTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	// Create a named tracer with package path as its name
	tracer := otel.Tracer("go.opentelemetry.io/contrib/examples/namedtracer")
	m0, _ := baggage.NewMemberRaw(string(fooKey), "foo1")
	m1, _ := baggage.NewMemberRaw(string(barKey), "bar1")
	b, _ := baggage.New(m0, m1)
	ctx = baggage.ContextWithBaggage(ctx, b)

	var span trace.Span
	ctx, span = tracer.Start(ctx, "operation")
	defer span.End()
	span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
	span.SetAttributes(anotherKey.String("yes"))
	if err := SubOperation(ctx); err != nil {
		panic(err)
	}

}
