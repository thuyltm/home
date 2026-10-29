FROM python:3.10-slim

RUN pip install dagster-pipes

COPY my_python_script.py .

ENTRYPOINT ["python", "my_python_script.py"]