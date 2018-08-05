[![Build Status](https://travis-ci.org/spatialcurrent/go-dfl.svg)](https://travis-ci.org/spatialcurrent/go-dfl) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-dfl)](https://goreportcard.com/report/spatialcurrent/go-dfl)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-dfl?status.svg)](https://godoc.org/github.com/spatialcurrent/go-dfl) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/go-spatial/tegola/blob/master/LICENSE.md)

# go-dfl

# Description

**go-dfl** is a Go implementation of the Dynamic Filter Language (DFL).

# Usage

**CLI**

You can also use the command line tool.

```
Usage: dfl -f INPUT [-verbose] [-version] [-help] [-env] [A=1] [B=2]
Options:
  -env
    	Load environment variables
  -f string
    	The DFL expression to evaulate
  -help
    	Print help
  -verbose
    	Provide verbose output
  -version
    	Prints version to stdout
```

**Go**

You can import **go-dfl** as a library with:

```
import (
  "github.com/spatialcurrent/go-dfl/dfl"
)
```

**JavaScript**

```html
<html>
  <head>
    <script src="https://...dfljs.js"></script>
  </head>
  <body>
    <script>
      var exp = "@pop + 10";
      var root = dfl.Parse(exp).Compile();
      var result = root.Evaluate({"pop": 10})
      ...
    </script>
  </body>
</html>
```

# Examples:

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

# Building

**CLI**

The command line DFL program can be built with the `scripts/build_cli.sh` script.

**JavaScript**

You can compile DFL to pure JavaScript with the `scripts/build_javascript.sh` script.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-dfl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
