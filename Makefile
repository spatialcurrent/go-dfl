# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

ifdef GOPATH
GCFLAGS=-trimpath=$(shell printenv GOPATH)/src
else
GCFLAGS=-trimpath=$(shell go env GOPATH)/src
endif

LDFLAGS=-X main.gitBranch=$(shell git branch | grep \* | cut -d ' ' -f2) -X main.gitCommit=$(shell git rev-list -1 HEAD)

ifndef DEST
DEST=bin
endif

.PHONY: help

help:  ## Print the help documentation
	@grep -E '^[a-zA-Z0-9_-\]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#
# Dependencies
#

deps_go:  ## Install Go dependencies
	go get -d -t ./...
	go get -d honnef.co/go/js/xhr # used in JavaScript build

.PHONY: deps_go_test
deps_go_test: ## Download Go dependencies for tests
	go get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # download shadow
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # install shadow
	go get -u github.com/kisielk/errcheck # download and install errcheck
	go get -u github.com/client9/misspell/cmd/misspell # download and install misspell
	go get -u github.com/gordonklaus/ineffassign # download and install ineffassign
	go get -u honnef.co/go/tools/cmd/staticcheck # download and instal staticcheck
	go get -u golang.org/x/tools/cmd/goimports # download and install goimports

deps_arm:  ## Install dependencies to cross-compile to ARM
	# ARMv7
	apt-get install -y libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev gcc-arm-linux-gnueabi g++-arm-linux-gnueabi
  # ARMv8
	apt-get install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu

deps_gopherjs:  ## Install GopherJS
	go get -u github.com/gopherjs/gopherjs

deps_javascript:  ## Install dependencies for JavaScript tests
	npm install .

#
# Go building, formatting, testing, and installing
#

.PHONY: fmt
fmt:  ## Format Go source code
	go fmt $$(go list ./... )

.PHONY: imports
imports: ## Update imports in Go source code
	# If missing, install goimports with: go get golang.org/x/tools/cmd/goimports
	goimports -w -local github.com/spatialcurrent/go-dfl,github.com/spatialcurrent/ $$(find . -iname '*.go')

.PHONY: vet
vet: ## Vet Go source code
	go vet $$(go list ./...)

.PHONY: test_go
test_go: ## Run Go tests
	bash scripts/test.sh

.PHONY: build
build: build_cli build_javascript build_so build_android  ## Build CLI, Shared Objects (.so), JavaScript, and Android

.PHONY: install
install:  ## Install dfl CLI on current platform
	go install -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/go-dfl/cmd/dfl

#
# Command line Programs
#

bin/dfl_darwin_amd64: ## Build dfl CLI for Darwin / amd64
	GOOS=darwin GOARCH=amd64 go build -o $(DEST)/dfl_darwin_amd64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/go-dfl/cmd/dfl

bin/dfl_linux_amd64: ## Build dfl CLI for Linux / amd64
	GOOS=linux GOARCH=amd64 go build -o $(DEST)/dfl_linux_amd64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/go-dfl/cmd/dfl

bin/dfl_windows_amd64.exe:  ## Build dfl CLI for Windows / amd64
	GOOS=windows GOARCH=amd64 go build -o $(DEST)/dfl_windows_amd64.exe -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/go-dfl/cmd/dfl

bin/dfl_linux_arm64: ## Build dfl CLI for Linux / arm64
	GOOS=linux GOARCH=arm64 go build -o $(DEST)/dfl_linux_arm64 -gcflags="$(GCFLAGS)" -ldflags="$(LDFLAGS)" github.com/spatialcurrent/go-dfl/cmd/dfl

.PHONY: build_cli
build_cli: bin/dfl_darwin_amd64 bin/dfl_linux_amd64 bin/dfl_windows_amd64.exe bin/dfl_linux_arm64  ## Build command line programs

#
# Shared Objects
#

bin/dfl.so:  ## Compile Shared Object for current platform
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	CGO_ENABLED=1 go build -o $(DEST)/dfl.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-dfl/plugins/dfl

bin/dfl_linux_amd64.so:  ## Compile Shared Object for Linux / amd64
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(DEST)/dfl_linux_amd64.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-dfl/plugins/dfl

bin/dfl_linux_armv7.so:  ## Compile Shared Object for Linux / ARMv7
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/dfl_linux_armv7.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-dfl/plugins/dfl

bin/dfl_linux_armv8.so:   ## Compile Shared Object for Linux / ARMv8
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	# Dependencies - https://www.96boards.org/blog/cross-compile-files-x86-linux-to-96boards/
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/dfl_linux_armv8.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-dfl/plugins/dfl

build_so: bin/dfl_linux_amd64.so bin/dfl_linux_armv7.so bin/dfl_linux_armv8.so  ## Build Shared Objects (.so)

#
# Android
#

bin/dfl.aar:  ## Build Android Archive Library
	gomobile bind -target android -javapkg=com.spatialcurrent -o $(DEST)/dfl.aar -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-dfl/pkg/dfl

build_android: bin/dfl.arr  ## Build artifacts for Android

#
# JavaScript
#

dist/dfl.mod.js:  ## Build JavaScript module
	gopherjs build -o dist/dfl.mod.js github.com/spatialcurrent/go-dfl/cmd/dfl.mod.js

dist/dfl.mod.min.js:  ## Build minified JavaScript module
	gopherjs build -m -o dist/dfl.mod.min.js github.com/spatialcurrent/go-dfl/cmd/dfl.mod.js

dist/dfl.global.js:  ## Build JavaScript library that attaches to global or window.
	gopherjs build -o dist/dfl.global.js github.com/spatialcurrent/go-dfl/cmd/dfl.global.js

dist/dfl.global.min.js:  ## Build minified JavaScript library that attaches to global or window.
	gopherjs build -m -o dist/dfl.global.min.js github.com/spatialcurrent/go-dfl/cmd/dfl.global.js

build_javascript: dist/dfl.mod.js dist/dfl.mod.min.js dist/dfl.global.js dist/dfl.global.min.js  ## Build artifacts for JavaScript

test_javascript:  ## Run JavaScript tests
	npm run test

lint:  ## Lint JavaScript source code
	npm run lint


#
# Examples
#

bin/dfl_example_c: bin/dfl.so  ## Build C example
	mkdir -p bin && cd bin && gcc -o dfl_example_c -I. ./../examples/c/main.c -L. -l:dfl.so

bin/dfl_example_cpp: bin/dfl.so  ## Build C++ example
	mkdir -p bin && cd bin && g++ -o dfl_example_cpp -I . ./../examples/cpp/main.cpp -L. -l:dfl.so

run_example_c: bin/dfl.so bin/dfl_example_c  ## Run C example
	cd bin && LD_LIBRARY_PATH=. ./dfl_example_c

run_example_cpp: bin/dfl.so bin/dfl_example_cpp  ## Run C++ example
	cd bin && LD_LIBRARY_PATH=. ./dfl_example_cpp

run_example_python: bin/dfl.so  ## Run Python example
	LD_LIBRARY_PATH=bin python examples/python/test.py

run_example_javascript: dist/dfl.mod.min.js  ## Run JavaScript module example
	node examples/js/index.mod.js

## Clean

clean:  ## Clean artifacts
	rm -fr bin
	rm -fr dist
