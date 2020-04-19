package didumeansub_test

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/subcommands"
	"github.com/mashiike/didumean"
	"github.com/mashiike/didumean/didumeansub"
)

type printCmd struct {
	capitalize bool
}

func (*printCmd) Name() string     { return "print" }
func (*printCmd) Synopsis() string { return "Print args to stdout." }
func (*printCmd) Usage() string {
	return `print [-capitalize] <some text>:
	Print args to stdout.
  `
}

func (p *printCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *printCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}

func ExampleExecute() {
	os.Args = []string{"./example", "print", "-captalize", "hoge"}
	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)
	flag.CommandLine.Init("./explaine", flag.ContinueOnError)
	subcommands.DefaultCommander.Error = &buf
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&printCmd{}, "")
	didumean.Parse()

	ctx := context.Background()
	didumeansub.Execute(ctx)

	fmt.Println(buf.String())
	// Output: flag provided but not defined: -captalize, did you mean -capitalize
	// print [-capitalize] <some text>:
	// 	Print args to stdout.
	//     -capitalize
	//     	capitalize output
}
