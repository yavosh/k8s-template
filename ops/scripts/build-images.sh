#!/usr/bin/env bash

docker buildx build -t "yavosh/k8s-test:latest" \
  --platform linux/amd64,linux/arm64,linux/arm/v7 --push .