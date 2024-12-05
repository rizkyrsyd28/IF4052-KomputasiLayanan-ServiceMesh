#!/bin/bash

docker build -t darktiger2280/service-consumer:latest .
docker image push darktiger2280/service-consumer:latest
docker rmi darktiger2280/service-consumer:latest || true