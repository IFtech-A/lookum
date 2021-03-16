#!/bin/sh

PREFIX=$1

DESCRIBE=$(git describe --match ${PREFIX}-* 2>/dev/null)
if [ -n "${DESCRIBE}" ]; then
    VERSION_FULL=$(echo ${DESCRIBE} | sed "s/^${PREFIX}-//")
    VERSION_MAJOR=$(echo ${VERSION_FULL} | awk -F- '{print $1}')
    VERSION_MINOR=$(echo ${VERSION_FULL} | sed "s/^${VERSION_MAJOR}//")
    if [ -n "${VERSION_MINOR}" ]; then
        VERSION_MINOR=$(echo ${VERSION_MINOR} | sed 's/^-/-p/')
    fi
else
    COMMITS=$(git rev-list --count HEAD)
    HASH=$(git rev-parse --short HEAD)
    VERSION_MAJOR=0.9.0
    VERSION_MINOR=-p${COMMITS}-g${HASH}
fi

echo ${VERSION_MAJOR}${VERSION_MINOR}