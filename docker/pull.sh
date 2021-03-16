#!/bin/sh
  
DIRNAME=$(dirname $0)
. ${DIRNAME}/docker.env
. ${DIRNAME}/docker-creds.env

for i in ${DOCKER_TYPES}; do
        (set -x; docker pull ${DOCKER_REG}/$i:latest)
done