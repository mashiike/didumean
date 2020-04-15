# didumean

![ci](https://github.com/mashiike/didumean/workflows/Test/badge.svg)
[![Documentation](https://godoc.org/github.com/mashiike/didumean?status.svg)](http://godoc.org/github.com/mashiike/didumean)
[![Go Report Card](https://goreportcard.com/badge/github.com/mashiike/didumean)](https://goreportcard.com/report/github.com/mashiike/didumean)


go flag package wrapper, for did you mean function

## Usage
see details in [godoc](http://godoc.org/github.com/mashiike/didumean)
```go
package main

import (
	"fmt"
	"flag"

	"github.com/mashiike/didumean"
)

func main() {  
    var (
        name string
    )
    flag.StringVar(&name, "name", "noname", "name parameter")
    didumean.Parse() // replace your flag.Parse()
    fmt.Println(name)
}
```

## LICENSE
MIT
