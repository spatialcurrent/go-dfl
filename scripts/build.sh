#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin
NAME=go-dfl

echo "******************"
echo "Formatting $DIR/../cmd/dfl"
cd $DIR/../cmd/dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building program for go-dfl"
cd $DIR/../bin
for GOOS in darwin linux windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "dfl_${GOOS}_amd64" github.com/spatialcurrent/go-dfl/cmd/dfl
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-dfl"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/dfl)"
