#!/usr/bin/env bash

echo "build image"
docker build -t sidecar:latest .

echo "down the container"
docker-compose down

echo "remove none images"
docker rmi `docker images | grep "<none>"` -f

echo "run image"
docker-compose up -d

echo "list container"
docker ps