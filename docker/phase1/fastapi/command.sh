#!/bin/bash
docker build -t thuyltm2201/my_fastapi_container_base:latest  ../../../ -f ./Dockerfile
docker push thuyltm2201/my_fastapi_container_base:latest
