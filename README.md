[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/fantasyname)](https://pkg.go.dev/github.com/s0rg/fantasyname)
[![Build](https://github.com/s0rg/fantasyname/workflows/ci/badge.svg)](https://github.com/s0rg/fantasyname/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/fantasyname)](https://goreportcard.com/report/github.com/s0rg/fantasyname)
[![Maintainability](https://api.codeclimate.com/v1/badges/95826f2eb57b89cefe70/maintainability)](https://codeclimate.com/github/s0rg/fantasyname/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/95826f2eb57b89cefe70/test_coverage)](https://codeclimate.com/github/s0rg/fantasyname/test_coverage)

[![License](https://img.shields.io/github/license/s0rg/fantasyname)](https://github.com/s0rg/fantasyname/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/fantasyname)](go.mod)


# fantasyname

This is a golang implementation of [name generator described at RinkWorks](http://rinkworks.com/namegen/)

# example

How its look like:
```go
    import "github.com/s0rg/fantasyname"

    gen, err := fantasyname.Compile("sV'i")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(gen.String())
```

[Here](https://github.com/s0rg/fantasyname/blob/master/_example/main.go) is a full example.

You can run it with `go run _example/main.go` to see results.
