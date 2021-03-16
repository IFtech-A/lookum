#!/bin/sh
  
usage()
{
        echo "Usage: $0 TARGET_DIR SRC_DIR"
        exit 1
}

if [ $# -ne 2 ]; then
        usage
fi

TARGET_DIR=$1
SRC_DIR=$2
echo "*********************************************************"

LINKS=$(find $TARGET_DIR -type l)
COUNT=$(echo $LINKS | wc -w)
YES=0
NO=0
for i in $LINKS; do
        SRC=$(readlink $i)
        case $SRC in
        $SRC_DIR*)
                rm -f $i
                cp -f $SRC $i
                YES=$((YES + 1))
                ;;
        *)
                NO=$((NO + 1))
                ;;
        esac
        echo -n "Translating links ... $YES/$NO/$COUNT\r"
done
echo "Translating links ... $YES/$NO/$COUNT done"
