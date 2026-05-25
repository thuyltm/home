# Use an official Python runtime as a parent image
FROM docker.io/python:3-slim

# Update the package list and install curl
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN pip install -r requirements.txt

RUN opentelemetry-bootstrap --action=install

# Make port 5000 available to the world outside this container
EXPOSE 5000

# Define environment variable
ENV OTEL_EXPORTER_OTLP_INSECURE="true"
ENV OTEL_LOGS_EXPORTER="otlp"
ENV OTEL_METRICS_EXPORTER="otlp"
ENV OTEL_TRACES_EXPORTER="otlp"
ENV OTEL_EXPORTER_OTLP_ENDPOINT="http://collector:4317"
ENV OTEL_SERVICE_NAME=flask-test

CMD ["opentelemetry-instrument", "flask", "--app", "rolldice_instrumentation_metric", "run", "--host=0.0.0.0"]