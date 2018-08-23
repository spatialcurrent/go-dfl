#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../dfl)"
cd $DIR/../dfl
go fmt
echo "Formatting $(realpath $DIR/../cmd/dfl)"
cd $DIR/../cmd/dfl
go fmt
echo "Done formatting."
echo "******************"
echo "Building program for go-dfl"
cd $DEST
for GOOS in darwin linux windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "dfl_${GOOS}_amd64" github.com/spatialcurrent/go-dfl/cmd/dfl
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-dfl"
    exit 1
fi
echo "Executables built at $DEST"
