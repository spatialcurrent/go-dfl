#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=${1:-$DIR/../bin}

mkdir -p $DEST

echo "******************"
echo "Building Javascript artifact for DFL"
cd $DEST
gopherjs build -o dfl.js github.com/spatialcurrent/go-dfl/cmd/dfl.js
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript for DFL"
    exit 1
fi
gopherjs build -m -o dfl.min.js github.com/spatialcurrent/go-dfl/cmd/dfl.js
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript artificats for DFL"
    exit 1
fi
echo "Executable built at $DEST"
