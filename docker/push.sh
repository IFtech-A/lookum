#!/bin/sh

DIRNAME=$(dirname $0)
. ${DIRNAME}/docker.env
. ${DIRNAME}/docker-creds.env

${DIRNAME}/login.sh

if [ $? -ne 0 ]; then
      exit 1;
fi

for i in ${DOCKER_TYPES}; do
      (set -x; docker push ${DOCKER_REG}/$i)
done

$DIRNAME/logout.sh