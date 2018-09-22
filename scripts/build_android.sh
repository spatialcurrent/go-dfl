#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=${1:-$DIR/../bin}

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../dfl)"
cd $DIR/../dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building AAR for dfl"
cd $DEST
gomobile bind -target android github.com/spatialcurrent/go-dfl/dfl
if [[ "$?" != 0 ]] ; then
    echo "Error building program for dfl"
    exit 1
fi
echo "Executable built at $DEST"
