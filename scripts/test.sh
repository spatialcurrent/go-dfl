#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
set -eu
echo "******************"
echo "Running unit tests"
cd $DIR/../dfl
go test
echo "******************"
echo "Using gometalinter with misspell, vet, ineffassign, and gosec"
echo "Testing $DIR/../dfl"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../dfl

echo "Testing $DIR/../plugins/dfl"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../plugins/dfl

echo "Testing $DIR/../cmd/dfl"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../cmd/dfl

echo "Testing $DIR/../cmd/dfl.js"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../cmd/dfl.js