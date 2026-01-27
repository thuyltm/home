1. Build and push my_fastapi_container_base
```sh
% cd docker/phase1/fastapi
% chmod a+x command.sh
% ./command.sh
```
2. Add _thuyltm2201/my_fastapi_container_base:latest_ oci image in MODULE.bazel
```sh
oci.pull(
    name="my_fastapi_container_base",
    digest="sha256:fecc54cdbca5a27eafd56abdd3518415e08239fb49dd0dfedc01ab6484079bed",
    image="docker.io/thuyltm2201/my_fastapi_container_base:latest",
    tag="latest",
)
```
3. To build the image and load it into it into a local runtime
```sh
docker run --rm -p 8000:8000 thuyltm2201/oci_python_hello_fastapi:latest
```