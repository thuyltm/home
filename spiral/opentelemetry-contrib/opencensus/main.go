package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ocmetric "go.opencensus.io/metric"
	"go.opencensus.io/metric/metricdata"
	"go.opencensus.io/metric/metricproducer"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	octrace "go.opencensus.io/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/bridge/opencensus"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	//instrumenttype differentiates between our gauge and view metrics
	keyType = tag.MustNewKey("instrumenttype")
	//Counts the number of lines read in form standard input
	countMeasure = stats.Int64("test_count", "A count of something", stats.UnitDimensionless)
	countView    = &view.View{
		Name:        "test_count",
		Measure:     countMeasure,
		Description: "A count of something",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{keyType},
	}
	serviceName = semconv.ServiceNameKey.String("test-service")
)

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

func main() {
	log.Println("Using OpenTelemetry stdout exporters")
	ctx := context.Background()
	conn, err := initConn()
	if err != nil {
		log.Fatal(err)
	}
	resName, err := resource.New(ctx,
		resource.WithAttributes(
			//The service name used to display traces in backends
			serviceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	//traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatal(fmt.Errorf("error creating trace exporter: %w", err))
	}
	metricsExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatal(fmt.Errorf("error creating metric exporter: %w", err))
	}
	tracing(ctx, resName, traceExporter)
	if err := monitoring(ctx, resName, metricsExporter); err != nil {
		log.Fatal(err)
	}
}

// tracing demonstrates overrudubg the OpenCensus
func monitoring(ctx context.Context, res *resource.Resource, exporter metric.Exporter) any {
	log.Println("Adding the OpenCensus metric Producer to an OpenTelemetry Reader to export OpenCensus metrics using the OpenTelemetry stdout exporter.")
	reader := metric.NewPeriodicReader(exporter, metric.WithProducer(opencensus.NewMetricProducer()))
	metric.NewMeterProvider(
		metric.WithReader(reader),
		metric.WithResource(res))
	log.Println("Registering a gauge metric using an OpenCensus registry")
	r := ocmetric.NewRegistry()
	metricproducer.GlobalManager().AddProducer(r)
	gauge, err := r.AddInt64Gauge(
		"test_gauge",
		ocmetric.WithDescription("A gauge for testing"),
		ocmetric.WithConstLabel(map[metricdata.LabelKey]metricdata.LabelValue{
			{Key: keyType.Name()}: metricdata.NewLabelValue("gauge"),
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to add gauge: %w", err)
	}
	entry, err := gauge.GetEntry()
	if err != nil {
		return fmt.Errorf("failed to get gauge entry: %w", err)
	}
	log.Println("Registering a cumulative metric using an OpenCensus view.")
	if err := view.Register(countView); err != nil {
		return fmt.Errorf("failed to register views: %w", err)
	}
	ctx, err = tag.New(ctx, tag.Insert(keyType, "view"))
	if err != nil {
		return fmt.Errorf("failed to set tag: %w", err)
	}
	for i := int64(1); true; i++ {
		entry.Set(i)
		stats.Record(ctx, countMeasure.M(1))
		time.Sleep(time.Second)
	}
	return nil
}

func tracing(ctx context.Context, res *resource.Resource, otExporter *otlptrace.Exporter) {
	log.Println("Configuring OpenCensus. Not Registering any OpenCensus exporters.")
	octrace.ApplyConfig(octrace.Config{DefaultSampler: octrace.AlwaysSample()})
	//using a batch span processor to aggregate spans before export
	bsp := sdktrace.NewBatchSpanProcessor(otExporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(otExporter),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	log.Println("Installing the OpenCensus bridge to make OpenCensus libraries write spans using OpenTelemetry.")
	opencensus.InstallTraceBridge()
	tp.ForceFlush(ctx)
	log.Println("Creating OpenCensus span, which should be printed out using the OpenTelemetry stdouttrace exporter.\n--It shoud have no parent, since it is the first span.")
	ctx, outerOCSpan1 := octrace.StartSpan(ctx, "OpenCensusOuterSpan")
	outerOCSpan1.End()
	tp.ForceFlush(ctx)
	log.Println("Creating OpenTelemetry span\n--It should have the OpenCensus span as a parent, since the OpenCensus span was written with using OpenTelemetry APIs.")
	ctx, outerOCSpan2 := octrace.StartSpan(ctx, "OpenCensusOuterSpan")
	outerOCSpan2.End()
	tp.ForceFlush(ctx)
	log.Println("Creating OpenTelemetry span\n--It should have the OpenCensus span as a parent, since the OpenCensus span was written with using OpenTelemetry APIs.")
	tracer := tp.Tracer("go.opentelemetry.io/contrib/examples/opencensus")
	ctx, otspan := tracer.Start(ctx, "OpenTelemetrySapn")
	otspan.End()
	tp.ForceFlush(ctx)
	log.Println("Creating OpenCensus span, which should be printed out using the OpenTelemetry stdoutrace exporter.\n--It should have the OpenTelemetry span as a parent, since it was written using OpenTelemetry APIs.")
	_, innerOCSpan := octrace.StartSpan(ctx, "OpenCensusInnerSpan")
	innerOCSpan.End()
	tp.ForceFlush(ctx)
}
