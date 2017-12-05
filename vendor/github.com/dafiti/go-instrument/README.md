# Golang Instrument

[![Build Status](https://img.shields.io/travis/dafiti/go-instrument/master.svg?style=flat-square)](https://travis-ci.org/dafiti/go-instrument)
[![Coverage Status](https://img.shields.io/coveralls/dafiti/go-instrument/master.svg?style=flat-square)](https://coveralls.io/github/dafiti/go-instrument?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/dafiti/go-instrument)

Simple interfaces for instrumentation

## Installation
```sh
go get github.com/dafiti/go-instrument
```

## Instruments
 - New Relic

## Usage Examples

```go
package main

import (
	"fmt"
	"github.com/dafiti/go-instrument"
	newrelic "github.com/newrelic/go-agent"
	"os"
)

var (
	inst instrument.Instrument
)

func main() {
	inst = new(instrument.NewRelic)

	app, err := newrelic.NewApplication(
		newrelic.NewConfig("Example App", os.Getenv("NEWRELIC_LICENSE_KEY")),
	)

	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	inst.SetTransaction(app.StartTransaction("background", nil, nil))

	fmt.Println(sum(10, 20))
}

func sum(v1 int, v2 int) int {
	defer inst.Segment("sum").End()

	return v1 + v2
}
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/dafiti/go-instrument](https://godoc.org/github.com/dafiti/go-instrument).

## License

This project is released under the MIT licence. See [LICENCE](https://github.com/dafiti/go-instrument/blob/master/LICENSE) for more details.
