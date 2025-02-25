package findgo_test

import (
	"testing"
	"testing/fstest"

	"github.com/ezebunandu/findgo"
	"github.com/google/go-cmp/cmp"
)

func TestFiles_CorrectlyListsFilesInMapFS(t *testing.T) {
    t.Parallel()
    fsys := fstest.MapFS{
        "file.go": {},
        "subfolder/subfolder.go": {},
        "subfolder2/another.go": {},
        "subfolder2/file.go": {},
    }
    want := []string{
        "file.go",
        "subfolder/subfolder.go",
        "subfolder2/another.go",
        "subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got) {
        t.Error(cmp.Diff(want, got))
    }
}