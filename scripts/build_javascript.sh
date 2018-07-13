#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $(realpath $DIR/../dfl)"
cd $DIR/../dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building Javascript for DFL"
cd $DIR/../bin
gopherjs build github.com/spatialcurrent/go-dfl/cmd/dfljs
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript for DFL"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin)"