#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=${1:-$DIR/../bin}

mkdir -p $DEST

echo "******************"
echo "Building Shared Object (*.so) for dfl"
cd $DEST
go build -o dfl.so -buildmode=c-shared github.com/spatialcurrent/go-dfl/plugins/dfl
if [[ "$?" != 0 ]] ; then
    echo "Error Building Shared Object (*.so) for dfl"
    exit 1
fi
echo "Shared Object (*.so) built at $DEST"
