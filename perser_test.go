package didumean_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/mashiike/didumean"
)

func TestFlagSet(t *testing.T) {
	cases := []struct {
		arguments []string
		err       string
		val       int
	}{
		{
			arguments: []string{"--huge", "1222", "-poyo", "1234", "dame"},
			err:       "flag provided but not defined: -huge, did you mean -hoge",
			val:       0,
		},
		{
			arguments: []string{"--hoge", "1222", "-poyo", "1234", "dame"},
			err:       "flag provided but not defined: -poyo, did you mean -piyo",
			val:       1222,
		},
		{
			arguments: []string{"--hoge", "1222", "-piyo", "1234", "dame"},
			err:       "",
			val:       1234,
		},
		{
			arguments: []string{"--hoge", "1222", "-piyo"},
			err:       "flag needs an argument: -piyo",
			val:       1222,
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.arguments, "_"), func(t *testing.T) {
			flagSet := didumean.NewFlagSet("testing", flag.ContinueOnError)
			var val int
			flagSet.IntVar(&val, "hoge", 0, "hogehoge")
			flagSet.IntVar(&val, "piyo", 0, "piyopiyo")
			flagSet.IntVar(&val, "fuga", 0, "fugafuga")

			err := flagSet.Parse(c.arguments)
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}
			if errStr != c.err {
				t.Logf("got      = %#v", errStr)
				t.Logf("expected = %#v", c.err)
				t.Error("unexpected error")
			}
			if val != c.val {
				t.Logf("got      = %#v", val)
				t.Logf("expected = %#v", c.val)
				t.Error("unexpected val")
			}
		})
	}
}
func TestParse(t *testing.T) {
	didumean.Parse()
}
