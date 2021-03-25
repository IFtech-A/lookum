#!/bin/sh

DIRNAME=$(dirname $0)

. ${DIRNAME}/docker.env


usage() {
    echo "Usage: $0 [api|db]";
    exit 1;
}

prepare() {
    DOCKER_TAG=$(${DIRNAME}/../script/git-version.sh)
}

build_image() {
    case ${DOCKER_TYPE} in
    api|db)
        (set -x; docker build \
            -t ${DOCKER_REG}/${DOCKER_TYPE}:${DOCKER_TAG} \
            -t ${DOCKER_REG}/${DOCKER_TYPE}:latest \
            -f ${DIRNAME}/dockerfiles/${DOCKER_TYPE}.Dockerfile \
            ${PACKAGE_BACKEND_DIR}/$DOCKER_TYPE
        )
    ;;
    proxy)
        (set -x; docker build \
            -t ${DOCKER_REG}/${DOCKER_TYPE}:${DOCKER_TAG} \
            -t ${DOCKER_REG}/${DOCKER_TYPE}:latest \
            -f ${DIRNAME}/dockerfiles/${DOCKER_TYPE}.Dockerfile \
            ${PACKAGE_PROXY_DIR}/$DOCKER_TYPE
        )
    ;;
    esac
}

case $1 in
api|db|proxy)
    prepare
    DOCKER_TYPE=$1 build_image
    ;;
all)
    prepare
    DOCKER_TYPE=api build_image
    DOCKER_TYPE=db build_image
    DOCKER_TYPE=proxy build_image
    ;;
*)
    usage
    ;;
esac