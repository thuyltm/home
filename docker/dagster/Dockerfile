FROM python:3.11

# Copy the uv executable from the official image
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

# Set environment variables to optimize for Docker environment
ENV UV_SYSTEM_PYTHON=1 \
    UV_COMPILE_BYTECODE=1

WORKDIR /dagster-postgresql-quickstart

COPY perimeter/dagster/dagster-postgresql-quickstart /dagster-postgresql-quickstart

# This make sure that logs show up immediately instead of being buffered
ENV PYTHONUNBUFFERED=1

RUN uv sync

ENV VIRTUAL_ENV=/dagster-postgresql-quickstart/.venv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

# Expose Dagster's default webserver port
EXPOSE 4000

#CMD ["dagster", "api", "grpc", "-h", "0.0.0.0", "-p", "4000", "-m", "my_grpc_dagster_project.definitions"]
