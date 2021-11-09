[![Build](https://github.com/s0rg/fantasyname/workflows/ci/badge.svg)](https://github.com/s0rg/fantasyname/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/fantasyname)](https://goreportcard.com/report/github.com/s0rg/fantasyname)
[![Maintainability](https://api.codeclimate.com/v1/badges/95826f2eb57b89cefe70/maintainability)](https://codeclimate.com/github/s0rg/fantasyname/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/95826f2eb57b89cefe70/test_coverage)](https://codeclimate.com/github/s0rg/fantasyname/test_coverage)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/fantasyname)](https://pkg.go.dev/github.com/s0rg/fantasyname)
[![License](https://img.shields.io/github/license/s0rg/fantasyname)](https://github.com/s0rg/fantasyname/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/fantasyname)](go.mod)


# fantasyname

This is a golang implementation of [name generator described at RinkWorks](http://rinkworks.com/namegen/)

# example

How it looks like:
```go
    import "github.com/s0rg/fantasyname"

    gen, err := fantasyname.Compile("sV'i")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(gen.String()) // will print something like: entheu'loaf
```

[Here](https://github.com/s0rg/fantasyname/blob/master/_example/main.go) is a full example.

You can run it with `go run _example/main.go` to see results.

# pattern syntax

The letters `s`, `v`, `V`, `c`, `B`, `C`, `i`, `m`, `M`, `D`, and `d` represent different types of random replacements:

 - `s` - generic syllable
 - `v` - vowel
 - `V` - vowel or vowel combination
 - `c` - consonant
 - `B` - consonant or consonant combination suitable for beginning a word
 - `C` - consonant or consonant combination suitable anywhere in a word
 - `i` - insult
 - `m` - mushy name
 - `M` - mushy name ending
 - `D` - consonant suited for a stupid person's name
 - `d` - syllable suited for a stupid person's name (begins with a vowel)

Everything else is emitted literally.

All characters between parenthesis `()` are emitted literally. For example, the pattern `s(dim)`,
emits a random generic syllable followed by `dim`.

Characters between angle brackets `<>` emit patterns from the table above.
Imagine the entire pattern is wrapped in one of these.

In both types of groupings, a vertical bar `|` denotes a random choice. Empty groups are allowed.
For example, `(foo|bar)` emits either `foo` or `bar`. The pattern `<c|v|>` emits a constant, vowel,
or nothing at all.

An exclamation point `!` means to capitalize the component that follows it. For example,
`!(foo)` will emit `Foo` and `v!s` will emit a lowercase vowel followed by a capitalized syllable, like `eRod`.
