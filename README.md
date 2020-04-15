# didumean

![ci](https://github.com/mashiike/urlio/workflows/Go/badge.svg)
[![Documentation](https://godoc.org/github.com/mashiike/urlio?status.svg)](http://godoc.org/github.com/mashiike/urlio)
[![Go Report Card](https://goreportcard.com/badge/github.com/mashiike/urlio)](https://goreportcard.com/report/github.com/mashiike/urlio)


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
