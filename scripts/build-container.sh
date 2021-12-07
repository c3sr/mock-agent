#!/usr/bin/env bash

cd `git rev-parse --show-toplevel`
source "scripts/functions.sh"

USAGE="Usage: ${0} <version> (registry)
    version       (required) The version number to tag the image with
                      Note: image will always be tagged as 'latest'
    registry      (optional) The Docker registry to tag images against
                      Defaults to 'mock-agent'"

VERSION="$1"
REGISTRY="$2"
: ${REGISTRY:="mock-agent"}
DOCKERFILE="./Dockerfile"

if [ -z $VERSION ]; then
    die_with_message "$USAGE"
fi

if [ ! -f "$DOCKERFILE" ]; then
    die_with_message "Dockerfile ${DOCKERFILE} does not exist"
fi

docker build -t "$REGISTRY/mock-agent:$VERSION" --file "$DOCKERFILE"
docker tag "$REGISTRY/mock-agent:$VERSION" "$REGISTRY/mock-agent:latest"

cleanup_and_exit