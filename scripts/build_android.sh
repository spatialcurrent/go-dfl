#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $(realpath $DIR/../dfl)"
cd $DIR/../dfl
go fmt
echo "Formatting $DIR/../dfl"
cd $DIR/../dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building AAR for dfl"
cd $DIR/../bin
gomobile bind -target android github.com/spatialcurrent/go-dfl/dfl
if [[ "$?" != 0 ]] ; then
    echo "Error building program for dfl"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin)"
