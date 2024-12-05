#!/bin/bash

docker build -t darktiger2280/service-tiket-b:latest .
docker image push darktiger2280/service-tiket-b:latest
docker rmi darktiger2280/service-tiket-b:latest || true