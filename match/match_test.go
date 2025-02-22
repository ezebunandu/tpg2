package match_test

import (
	"testing"
    "os"

	"github.com/ezebunandu/match"
    "github.com/rogpeppe/go-internal/testscript"
)

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"match": match.Main,
	}))
}