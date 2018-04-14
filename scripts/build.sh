#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin
NAME=go-dfl

echo "******************"
echo "Formatting $DIR/../cmd/dfl"
cd $DIR/../cmd/dfl
go fmt
echo "Formatting github.com/spatialcurrent/$NAME/dfl"
go fmt github.com/spatialcurrent/$NAME/dfl
echo "Done formatting."
echo "******************"
echo "Building program for $NAME"
cd $DIR/../bin
go build github.com/spatialcurrent/$NAME/cmd/dfl
if [[ "$?" != 0 ]] ; then
    echo "Error building program for $NAME"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/dfl)"
