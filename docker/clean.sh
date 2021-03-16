#!/bin/sh
  
DIRNAME=$(dirname $0)
. $DIRNAME/docker.env
. $DIRNAME/docker-creds.env

images()
{
        docker images $DOCKER_REG/* -a "$@"
}

remove_danglings()
{
        IMAGES=$(images -q -f dangling=true | sort -u)
        if [ -n "$IMAGES" ]; then
                (set -x; docker rmi $IMAGES)
        else
                echo "% No dangling images"
        fi
}

remove_by_tags()
{
        TAGS=$(images | awk '{print $2}' | sort -uV)
        for t in $TAGS; do
                case $t in
                latest|TAG)
                        ;;
                *)
                        echo $t
                        ;;
                esac
        done

        read -p '> Input tags to remove: ' INPUT
        for i in $INPUT; do
                IMAGES=$(docker images --format '{{.Repository}}:{{.Tag}}' $DOCKER_REG/*:$i)
                if [ -n "$IMAGES" ]; then
                        (set -x; docker rmi $IMAGES)
                fi
        done
}

remove_danglings
echo
remove_by_tags