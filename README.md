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

# todo

[] more test cover
