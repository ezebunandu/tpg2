package older_test

import (
    "testing"
    "testing/fstest"
    "time"
    "github.com/google/go-cmp/cmp"
    "github.com/ezebunandu/older"
)

func TestOlderFiles_ReturnsFilesOlderThanGivenDuration(t *testing.T) {
    t.Parallel()
    now := time.Now()
    fsys := fstest.MapFS{
        "file.go": {ModTime: now.AddDate(0,0,-29)},
        "subfolder/subfolder.go": {ModTime: now.AddDate(0,0,-35)},
        "subfolder2/another.go": {ModTime: now.AddDate(0,0,-23)},
        "subfolder2/file.go": {ModTime: now.AddDate(0,0,-45)},
    }
    want := []string{

        "subfolder/subfolder.go",
        "subfolder2/file.go",
    }
    got := older.OlderFiles(fsys, 30 * time.Hour*24)
    if !cmp.Equal(want, got) {
        t.Error(cmp.Diff(want, got))
    }
}