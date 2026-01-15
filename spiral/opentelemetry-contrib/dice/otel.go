package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceName = semconv.ServiceNameKey.String("dice-service")

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup
func setupOTelSDK(ctx context.Context) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error
	var err error
	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined
	// Each registered cleanup will be invoked once
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
	//handleErr calls shutdown for cleanup and makes sure that all errors are returned
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	// setup propagator
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	//set up trace provider
	conn, err := initConn()
	if err != nil {
		panic(err)
	}
	res, err := resource.New(ctx,
		resource.WithAttributes(
			//The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		panic(err)
	}
	tracerProvider, err := newTracerProvider(ctx, res, conn)
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	// setup meter provider
	meterProvider, err := newMeterProvider(ctx, res, conn)
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)
	// set up logger provider
	loggerProvider, err := newLoggerProvider(ctx, res, conn)
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)
	return shutdown, err
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
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

func newTracerProvider(ctx context.Context,
	res *resource.Resource, conn *grpc.ClientConn) (*trace.TracerProvider, error) {
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}
	bsp := trace.WithBatcher(traceExporter, trace.WithBatchTimeout(5*time.Second))
	tracerProvider := trace.NewTracerProvider(
		bsp,
		trace.WithResource(res),
	)
	return tracerProvider, nil
}

func newMeterProvider(ctx context.Context,
	res *resource.Resource, conn *grpc.ClientConn) (*metric.MeterProvider, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)
	return meterProvider, nil
}

func newLoggerProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (*log.LoggerProvider, error) {
	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
		log.WithResource(res),
	)
	return loggerProvider, nil
}
