[![GoDoc](https://godoc.org/github.com/sprungknoedl/dcfg?status.svg)](https://godoc.org/github.com/sprungknoedl/dcfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sprungknoedl/dcfg)](https://goreportcard.com/report/github.com/sprungknoedl/dcfg)
[![Build Status](https://img.shields.io/travis/sprungknoedl/dcfg.svg)](https://travis-ci.org/sprungknoedl/dcfg)

# dcfg
dcfg retrieves configuration values for twelve-factor apps. In addition to
values from the environment, this package can also retrieve values from docker
secret files and set default values.

## Installation
```
go get github.com/sprungknoedl/dcfg
```

## Usage
```go
package main

import "github.com/sprungknoedl/dcfg"

func main() {
    // optional: set defaults
    dcfg.Defaults.Set("HOST", "localhost")
    dcfg.Defaults.Set("PORT", "8080")

    // ignore errors for example, don't do this in production ;)
    host, _ := dcfg.GetString("HOST")
    port, _ := dcfg.GetInt("PORT")
}
```
