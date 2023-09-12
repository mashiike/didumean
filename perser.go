package didumean

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/agnivade/levenshtein"
)

// FlagSet is *flag.FlagSet wrapper with did you mean function.
type FlagSet struct {
	*flag.FlagSet

	errorHandling flag.ErrorHandling
	output        io.Writer
}

// NewFlagSet create wraped FlagSet
func NewFlagSet(name string, handling flag.ErrorHandling) *FlagSet {
	return wrapFlagSet(flag.NewFlagSet(name, handling))
}

func wrapFlagSet(orig *flag.FlagSet) *FlagSet {
	return &FlagSet{
		FlagSet: orig,
	}
}

const (
	trapErrorString = "flag provided but not defined:"
	threshold       = 3
)

// Parse as flag.FlagSet.Parse with did you mean.
func (f *FlagSet) Parse(arguments []string) error {
	f.errorHandling = f.FlagSet.ErrorHandling()
	f.FlagSet.Init(f.FlagSet.Name(), flag.ContinueOnError)

	err := f.parse(arguments)
	if err == nil {
		return nil
	}
	switch f.errorHandling {
	case flag.ContinueOnError:
		return err
	case flag.ExitOnError:
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	case flag.PanicOnError:
		panic(err)
	}
	return nil
}

func (f *FlagSet) parse(arguments []string) error {
	var buf bytes.Buffer
	f.output = f.FlagSet.Output()
	f.FlagSet.SetOutput(&buf)

	err := f.FlagSet.Parse(arguments)
	msg := buf.String()
	defer func() {
		if f.output != nil {
			io.WriteString(f.output, msg)
		}
		f.FlagSet.SetOutput(f.output)
	}()
	if err == nil {
		return nil
	}

	if !strings.HasPrefix(err.Error(), trapErrorString) {
		return err
	}
	invalidFlag := arguments[len(arguments)-f.FlagSet.NArg()-1]
	invalidFlag = strings.TrimPrefix(strings.TrimPrefix(invalidFlag, "-"), "-")
	minDistance := len(invalidFlag) + 1
	minFlag := ""
	f.FlagSet.VisitAll(func(fl *flag.Flag) {
		distance := levenshtein.ComputeDistance(invalidFlag, fl.Name)
		if distance < threshold && distance < minDistance {
			minDistance = distance
			minFlag = fl.Name
		}
	})
	if minFlag != "" {
		newErr := fmt.Errorf("%s, did you mean -%s", err, minFlag)
		msg = strings.Replace(msg, err.Error(), newErr.Error(), 1)
		return newErr
	}

	return err
}

// Parse as flag.Parse with did you mean.
func Parse() {
	set := wrapFlagSet(flag.CommandLine)
	set.Parse(os.Args[1:])
}
