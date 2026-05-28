
from random import randint
from flask import Flask, request
import logging
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    BatchSpanProcessor
)
from opentelemetry import metrics
from opentelemetry.exporter.otlp.proto.http.metric_exporter import OTLPMetricExporter
from opentelemetry.sdk.metrics import MeterProvider
from opentelemetry.sdk.metrics.export import (
    PeriodicExportingMetricReader,
)
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.sdk._logs.export import BatchLogRecordProcessor, ConsoleLogRecordExporter
from opentelemetry._logs import set_logger_provider

from opentelemetry.sdk.resources import SERVICE_NAME, Resource

# Service name is required for most backends
resource = Resource.create(attributes={
    SERVICE_NAME: "flask-test"
})

tracerProvider = TracerProvider(resource=resource)
processor = BatchSpanProcessor(OTLPSpanExporter(endpoint="http://otelcol:4318/v1/traces"))
tracerProvider.add_span_processor(processor)

# Sets the global default tracer provider
trace.set_tracer_provider(tracerProvider)
# Creates a tracer from the global tracer provider
tracer = trace.get_tracer("diceroller.tracer")

metric_reader = PeriodicExportingMetricReader(
    OTLPMetricExporter(endpoint="http://otelcol:4318/v1/metrics")
)
meterProvider = MeterProvider(resource=resource, metric_readers=[metric_reader])

# Sets the global default meter provider
metrics.set_meter_provider(meterProvider)
# Creates a meter from the global meter provider
meter = metrics.get_meter("diceroller.meter")

roll_counter = meter.create_counter(
    "dice.rolls",
    description="The number of rolls by roll value"
)

app = Flask(__name__)
logProvider = LoggerProvider()
logProcessor = BatchLogRecordProcessor(ConsoleLogRecordExporter())
logProvider.add_log_record_processor(logProcessor)
# Sets the global default logger provider
set_logger_provider(logProvider)
handler = LoggingHandler(level=logging.INFO, logger_provider=logProvider)
logging.basicConfig(handlers=[handler], level=logging.INFO)
logger = logging.getLogger(__name__)


@app.route("/flask-test/rolldice")
def roll_dice():
    with tracer.start_as_current_span("roll") as roll_span:
        player = request.args.get('player', default = None, type = str)
        result = str(roll())
        roll_span.set_attribute("roll.value", result)
        roll_counter.add(1, {"roll.value": result})
        if player:
            logger.warning("%s is rolling the dice: %s", player, result)
        else:
            logger.warning("Anonymous player is rolling the dice: %s", result)
        return result

def roll():
    return randint(1, 6)

@app.route('/flask-test/health', methods=['GET'])
def health_check():
    return {"status": "healthy"}, 200

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000, debug=True)