package didumean_test

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/mashiike/didumean"
)

func ExampleParse() {
	var buf bytes.Buffer
	flag.CommandLine = flag.NewFlagSet("example", flag.ContinueOnError)
	flag.CommandLine.SetOutput(&buf)
	os.Args = []string{"example", "-staing", "hoge"}

	var (
		val int
		str string
		on  bool
	)
	flag.IntVar(&val, "value", 0, "value")
	flag.StringVar(&str, "string", "", "string")
	flag.BoolVar(&on, "on", false, "on")
	didumean.Parse()

	fmt.Println(buf.String())
	// Output: flag provided but not defined: -staing, did you mean -string
	// Usage of example:
	//   -on
	//     	on
	//   -string string
	//     	string
	//   -value int
	//     	value
}
