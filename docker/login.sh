#!/bin/sh
  
DIRNAME=$(dirname $0)
. $DIRNAME/docker.env
. $DIRNAME/docker-creds.env

DOCKER_OPT=
if [ -n "${DOCKER_USERNAME}" ]; then
        DOCKER_OPT="-u ${DOCKER_USERNAME}"
fi

(set -x; docker login ${DOCKER_OPT})