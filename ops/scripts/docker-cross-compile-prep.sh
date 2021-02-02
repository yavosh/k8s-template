#!/usr/bin/env bash

## Needs experimental && cli experimental

docker buildx create --name mybuilder
docker buildx use mybuilder
docker buildx inspect --bootstrap