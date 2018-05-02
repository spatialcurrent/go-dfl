[![Build Status](https://travis-ci.org/spatialcurrent/go-dfl.svg)](https://travis-ci.org/spatialcurrent/go-dfl) [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-dfl?status.svg)](https://godoc.org/github.com/spatialcurrent/go-dfl)

# go-dfl

# Description

**go-dfl** is a Go implementation of the Dynamic Filter Language (DFL).

# Usage

You can import **go-dfl** as a library with:

```
import (
  "github.com/spatialcurrent/go-dfl/dfl"
)
```

You can also test DFL filtering using a command line tool:

```
Usage: dfl -filter INPUT [-verbose] [-version] [-help] [A=1] [B=2]
Options:
  -filter string
    	The DFL expression to evaulate
  -help
    	Print help
  -verbose
    	Provide verbose output
  -version
    	Prints version to stdout

```

# Examples:

```
./dfl -verbose -filter '@pop > (10 - 2)' craft=brewery name=Stone pop=10
true
Done in 7.354125ms
```

```
./dfl -verbose -filter '@craft like "brewery"' craft=brewery name=Stone pop=10
true
Done in 3.44965ms
```

```
./dfl -verbose -filter '@craft like "brew%"' craft=brewery name=Stone pop=10
true
Done in 4.159278ms
```

```
./dfl -verbose -filter '@craft ilike "Stone%"' craft=brewery name=Atlas pop=10
false
Done in 4.439012ms
```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-dfl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
