[![CircleCI](https://circleci.com/gh/spatialcurrent/go-dfl/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-dfl/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-dfl)](https://goreportcard.com/report/spatialcurrent/go-dfl)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-dfl?status.svg)](https://godoc.org/github.com/spatialcurrent/go-dfl) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-dfl/blob/master/LICENSE.md)

# go-dfl

## Description

**go-dfl** is a Go implementation of the Dynamic Filter Language (DFL).  **go-dfl** depends on:
- [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions) for many of the basic functions.
- [go-counter](https://github.com/spatialcurrent/go-reader-writer) for reading and writer files.
- [go-counter](https://github.com/spatialcurrent/go-counter) for counting for statistical functions.
- [go-try-get](https://github.com/spatialcurrent/go-try-get) for abstracting how to get values by name from unknown objects.

Using cross compilers, this library can also be called by other languages.  This library is cross compiled into a Shared Object file (`*.so`).  The Shared Object file can be called by `C`, `C++`, and `Python` on Linux machines.  See the examples folder for patterns that you can use.  This library is also compiled to pure `JavaScript` using [GopherJS](https://github.com/gopherjs/gopherjs).

## Usage

**CLI**

The command line tool, `dfl`, can be used to easily covert data between formats.  We currently support the following platforms.

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

Pull requests to support other platforms are welcome!  See the [examples](#examples) section below for usage.

**Go**

You can install the go-dfl packages with.


```shell
go get -u -d github.com/spatialcurrent/go-dfl/...
```

You can then import the `dfl` package with:

```go
import (
  "github.com/spatialcurrent/go-dfl/pkg/dfl"
)
```

See [dfl](https://godoc.org/github.com/spatialcurrent/go-dfl/pkg/dfl) in GoDoc for information on how to use Go API.

**Node**

DFL is built as a module.  In modern JavaScript, the module can be imported using [destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment).

```javascript
const { parse, compile, evaluate } = require('./dist/dfl.mod.min.js');
```

In legacy JavaScript, you can use the `dfl.global.js` file that simply adds `dfl` to the global scope.

**C**

A variant of the reader and writer functions are exported in a Shared Object file (`*.so`), which can be called by `C`, `C++`, and `Python` programs on Linux machines.  For complete patterns for `C`, `C++`, and `Python`, see the `examples` folder in this repo.

## Releases

**go-dfl** is currently in **alpha**.  See releases at https://github.com/spatialcurrent/go-dfl/releases.  See the **Building** section below to build from scratch.

**JavaScript**

- `dfl.global.js`, `dfl.global.js.map` - JavaScript global build  with source map
- `dfl.global.min.js`, `dfl.global.min.js.map` - Minified JavaScript global build with source map
- `dfl.mod.js`, `dfl.mod.js.map` - JavaScript module build  with source map
- `dfl.mod.min.js`, `dfl.mod.min.js.map` - Minified JavaScript module with source map

**Darwin**

- `dfl_darwin_amd64` - CLI for Darwin on amd64 (includes `macOS` and `iOS` platforms)

**Linux**

- `dfl_linux_amd64` - CLI for Linux on amd64
- `dfl_linux_amd64` - CLI for Linux on arm64
- `dfl_linux_amd64.h`, `dfl_linuxamd64.so` - Shared Object for Linux on amd64
- `dfl_linux_armv7.h`, `dfl_linux_armv7.so` - Shared Object for Linux on ARMv7
- `dfl_linux_armv8.h`, `dfl_linux_armv8.so` - Shared Object for Linux on ARMv8

**Windows**

- `dfl_windows_amd64.exe` - CLI for Windows on amd64

## Examples:

### Go

See the examples in the [dfl](https://godoc.org/github.com/spatialcurrent/go-dfl/pkg/dfl) package documentation.

### CLI

**Environment**

With the `-env` flag you can use DFL filters against the current environment.

```
./dfl -env -f '@SHELL in [/bin/sh, /bin/bash]' && echo "Shell is set to sh or bash"
```

**Arrays**

You can use DFL to filter integer arrays.

```
./dfl -f '@a == [1, 2, 3, 4]' 'a=[1, 2, 3, 4]'
# returns true as exit code 0
```

You can also use DFl to compare byte arrays.

```
./dfl -f '@a == [137, 80, 78, 71]' 'a=[0x89, 0x50, 0x4E, 0X47]'
# returns true as exit code 0
```

**IP Address Information**

You can use DFL to filter IP addresses.

```
./dfl -f '@ip in 10.10.0.0/16' ip=10.10.20.22
# returns true as exit code 0
```

**OpenStreetMap**

You can also use DFL to filter OpenStreetMap features.

```
./dfl -f '@pop > (10 - 2)' craft=brewery name=Stone pop=10
# returns true as exit code 0
```

```
./dfl -f '@craft like "brewery"' craft=brewery name=Stone pop=10
# returns true as exit code 0
```

```
./dfl -f '@craft like "brew%"' craft=brewery name=Stone pop=10
# returns true as exit code 0
```

```
./dfl -f '@craft ilike "Stone%"' craft=brewery name=Atlas pop=10
# returns true as exit code 0
```

## Building

Use `make help` to see help information for each target.

**CLI**

The `make build_cli` script is used to build executables for Linux and Windows.

**JavaScript**

You can compile DFL to pure JavaScript with the `make build_javascript` script.

**Shared Object**

The `make build_so` script is used to build Shared Objects (`*.so`), which can be called by `C`, `C++`, and `Python` on Linux machines.

**Changing Destination**

The default destination for build artifacts is `go-dfl/bin`, but you can change the destination with an environment variable.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `DEST=/usr/local/go/bin make build_cli`

## Testing

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

**JavaScript**

To run JavaScript tests, first install [Jest](https://jestjs.io/) using `make deps_javascript`, use [Yarn](https://yarnpkg.com/en/), or another method.  Then, build the JavaScript module with `make build_javascript`.  To run tests, use `make test_javascript`.  You can also use the scripts in the `package.json`.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-dfl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
