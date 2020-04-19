package didumeansub

import (
	"context"
	"flag"
	"io"

	"github.com/google/subcommands"
	"github.com/mashiike/didumean"
)

type Commander struct {
	*subcommands.Commander
	topFlags *flag.FlagSet
}

func NewCommander(topLevelFlags *flag.FlagSet, name string) *Commander {
	cdr := subcommands.NewCommander(topLevelFlags, name)
	return wrapCommander(topLevelFlags, cdr)

}

func wrapCommander(topLevelFlags *flag.FlagSet, cdr *subcommands.Commander) *Commander {
	return &Commander{
		Commander: cdr,
		topFlags:  topLevelFlags,
	}
}

func (cdr *Commander) Execute(ctx context.Context, args ...interface{}) subcommands.ExitStatus {
	if cdr.topFlags.NArg() < 1 {
		cdr.topFlags.Usage()
		return subcommands.ExitUsageError
	}

	name := cdr.topFlags.Arg(0)
	var targetCmd subcommands.Command
	cdr.Commander.VisitCommands(func(_ *subcommands.CommandGroup, cmd subcommands.Command) {
		if name == cmd.Name() {
			targetCmd = cmd
		}
	})

	if targetCmd == nil {
		cdr.topFlags.Usage()
		return subcommands.ExitUsageError
	}

	f := didumean.NewFlagSet(name, flag.ContinueOnError)
	f.SetOutput(cdr.Error)
	f.Usage = func() {
		io.WriteString(f.Output(), targetCmd.Usage())
		f.PrintDefaults()
	}
	targetCmd.SetFlags(f.FlagSet)
	if f.Parse(cdr.topFlags.Args()[1:]) != nil {
		return subcommands.ExitUsageError
	}
	return targetCmd.Execute(ctx, f.FlagSet, args...)
}

func Execute(ctx context.Context, args ...interface{}) subcommands.ExitStatus {
	wraped := wrapCommander(flag.CommandLine, subcommands.DefaultCommander)
	return wraped.Execute(ctx, args...)
}
