#!/bin/bash
bazel run //elliptic/serviceB:load

# OR another build uses ko
#export KO_DOCKER_REPO=thuyltm2201
#ko build .
echo $IMAGE
docker push $IMAGE