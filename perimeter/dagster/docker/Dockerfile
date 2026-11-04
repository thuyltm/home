FROM python:3.11-slim


# 2. Performance & Log Tuning
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on

WORKDIR /dagster-postgresql-quickstart

# 4. Copy the application code (ignores .venv thanks to .dockerignore)
COPY perimeter/dagster/dagster-postgresql-quickstart .

# 2. Build the environment inside the container
RUN python -m venv . && \
    ./bin/pip install --no-cache-dir --upgrade pip && \
    ./bin/pip install --no-cache-dir -r requirements.txt

# 3. Add the container's venv to the PATH so you don't have to call it explicitly
ENV PATH="/dagster-postgresql-quickstart/bin:$PATH"

ENV PYTHONPATH=/dagster-postgresql-quickstart/src

# Expose Dagster's default webserver port
EXPOSE 4000

#CMD ["dagster", "api", "grpc", "-h", "0.0.0.0", "-p", "4000", "-m", "dagster_postgresql_quickstart.definitions"]
