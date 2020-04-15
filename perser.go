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

type FlagSet struct {
	orig          *flag.FlagSet
	errorHandling flag.ErrorHandling
	output        io.Writer
}

func NewFlagSet(orig *flag.FlagSet) *FlagSet {
	return &FlagSet{
		orig: orig,
	}
}

const (
	trapErrorString = "flag provided but not defined:"
	threshold       = 3
)

//Parse as flag.FlagSet.Parse with did you mean.
func (f *FlagSet) Parse(arguments []string) error {
	f.errorHandling = f.orig.ErrorHandling()
	f.orig.Init(f.orig.Name(), flag.ContinueOnError)

	err := f.parse(arguments)

	switch f.errorHandling {
	case flag.ContinueOnError:
		return err
	case flag.ExitOnError:
		os.Exit(2)
	case flag.PanicOnError:
		panic(err)
	}
	return nil
}

func (f *FlagSet) parse(arguments []string) error {
	var buf bytes.Buffer
	f.output = f.orig.Output()
	f.orig.SetOutput(&buf)

	err := f.orig.Parse(arguments)
	msg := buf.String()
	defer func() {
		if f.output != nil {
			io.WriteString(f.output, msg)
		}
	}()
	if err == nil {
		return nil
	}

	if !strings.HasPrefix(err.Error(), trapErrorString) {
		return err
	}
	invalidFlag := arguments[len(arguments)-f.orig.NArg()-1]
	invalidFlag = strings.TrimPrefix(strings.TrimPrefix(invalidFlag, "-"), "-")
	minDistance := len(invalidFlag) + 1
	minFlag := ""
	f.orig.VisitAll(func(fl *flag.Flag) {
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

//Parse as flag.Parse with did you mean.
func Parse() {
	set := NewFlagSet(flag.CommandLine)
	set.Parse(os.Args[1:])
}
