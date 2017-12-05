# Echo Middlewares

[![Build Status](https://img.shields.io/travis/dafiti/echo-middleware/master.svg?style=flat-square)](https://travis-ci.org/dafiti/echo-middleware)
[![Coverage Status](https://img.shields.io/coveralls/dafiti/echo-middleware/master.svg?style=flat-square)](https://coveralls.io/github/dafiti/echo-middleware?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/dafiti/echo-middleware)

Middlewares for Echo Framework

## Installation
```sh
go get github.com/dafiti/echo-middleware
```

## Middlewares
 - New Relic

## Usage Examples

```go
package main

import (
    "net/http"
	"github.com/dafiti/echo-middleware"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.Use(middleware.NewRelic("app name", "license key"))

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/dafiti/echo-middleware](https://godoc.org/github.com/dafiti/echo-middleware).

## License

This project is released under the MIT licence. See [LICENCE](https://github.com/dafiti/echo-middleware/blob/master/LICENSE) for more details.
