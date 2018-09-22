#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "******************"
echo "Formatting $DIR/../dfl"
cd $DIR/../dfl
go fmt
echo "Formatting $DIR/../dfljs"
cd $DIR/../dfljs
go fmt
echo "Formatting $DIR/../plugins/dfl"
cd $DIR/../plugins/dfl
go fmt
echo "Formatting $DIR/../cmd/dfl"
cd $DIR/../cmd/dfl/
go fmt
echo "Formatting $DIR/../cmd/dfl.js"
cd $DIR/../cmd/dfl.js
go fmt