#!/bin/bash

docker build -t darktiger2280/service-tiket-a:latest .
docker image push darktiger2280/service-tiket-a:latest
docker rmi darktiger2280/service-tiket-a:latest || true