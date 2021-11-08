# fantasyname

This is a golang implementation of [name generator described at RinkWorks](http://rinkworks.com/namegen/)

# example
```go
    gen, err := fantasyname.Parse("sV'i")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(gen.String())
```

# todo

[] test cover

[] better readme
