#!/bin/bash
bazel run //elliptic/serviceB:load

# OR another build uses ko
#echo "Building image $IMAGE using a custom script..."
#export KO_DOCKER_REPO=thuyltm2201
#ko build -B --sbom=none .
echo $IMAGE
docker push $IMAGE