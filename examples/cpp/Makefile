# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build_so:
	go build -o dfl.so -buildmode=c-shared github.com/spatialcurrent/go-dfl/plugins/dfl
build_cpp:
	mkdir -p bin && g++ -o bin/dfl_test_cpp test.cpp -L. -l:dfl.so
build: build_so build_cpp
run:
	LD_LIBRARY_PATH=. bin/dfl_test_cpp
clean:
	rm -f dfl.h dfl.so bin/dfl_test_cpp
