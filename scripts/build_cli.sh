#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

set -eu

DEST=${1:-$DIR/../bin}

mkdir -p $DEST

echo "******************"
echo "Building program for go-dfl"
cd $DEST
for GOOS in linux darwin windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "dfl_${GOOS}_amd64" github.com/spatialcurrent/go-dfl/cmd/dfl
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-dfl"
    exit 1
fi
echo "Executables built at $DEST"
