#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $DIR/dfl"
cd $DIR/../dfl
go fmt
echo "Formatting $DIR/../cmd/dfl"
cd $DIR/../cmd/dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building Shared Object (*.so) for dfl"
cd $DIR/../bin
go build -o dfl.so -buildmode=c-shared github.com/spatialcurrent/go-dfl/plugins/dfl
if [[ "$?" != 0 ]] ; then
    echo "Error Building Shared Object (*.so) for dfl"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin)"
