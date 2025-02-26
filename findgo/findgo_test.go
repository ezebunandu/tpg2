package findgo_test

import (
	"archive/zip"
	"os"
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

func TestFiles_CorrectlyLIstsFilesInTree(t *testing.T){
    t.Parallel()
    fsys := os.DirFS("testdata/tree")
    want := []string{
        "file.go",
        "subfolder/subfolder.go",
        "subfolder2/another.go",
        "subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestFiles_CorrectlyLIstsFilesInZIPArchive(t *testing.T){
    t.Parallel()
    fsys, err := zip.OpenReader("testdata/files.zip")
    if err != nil {
        t.Fatal(err)
    }
    want := []string{
        "tree/file.go",
        "tree/subfolder/subfolder.go",
        "tree/subfolder2/another.go",
        "tree/subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func BenchmarkFilesOnDisk(b *testing.B){
    fsys := os.DirFS("testdata/tree")
    b.ResetTimer()
    for range b.N {
        _ = findgo.Files(fsys)
    }
}

func BenchmarkFilesInMapFS(b *testing.B){
    fsys := fstest.MapFS{
        "file.go": {},
        "subfolder/subfolder.go": {},
        "subfolder2/another.go": {},
        "subfolder2/file.go": {},
    }
    b.ResetTimer()
    for range b.N {
        _ = findgo.Files(fsys)
    }
}