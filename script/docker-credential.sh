#!/bin/sh
  
usage()
{
        echo "Usage: $0 CREDENTIAL_FILE"
        exit 1
}

if [ $# -ne 1 ]; then
        usage
fi

CREDENTIAL_FILE=$1
if [ -r $CREDENTIAL_FILE ]; then
        . $CREDENTIAL_FILE
fi

if [ -z "$DOCKER_USERNAME" ]; then
        read -p "Docker Username: " DOCKER_USERNAME
        echo DOCKER_USERNAME=${DOCKER_USERNAME} > $CREDENTIAL_FILE
fi
