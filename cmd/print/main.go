package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mashiike/didumean"
)

func main() {
	var (
		capitalize bool
	)
	flag.BoolVar(&capitalize, "capitalize", false, "capitalize output")
	didumean.Parse()
	for _, arg := range flag.Args() {
		if capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
}
