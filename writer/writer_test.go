package writer_test

import (
    "os"
    "testing"
    "github.com/rogpeppe/go-internal/testscript"
    "github.com/ezebunandu/writer"
    "github.com/google/go-cmp/cmp"
)

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"writefile": writer.Main,
	}))
}

func TestWriteToFile_WritesGivenDataToFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/write_test.txt"
    want := []byte{1, 2, 3}

    err := writer.WriteToFile(path, want)
    if err != nil {
        t.Fatal(err)
    }

    stat, err := os.Stat(path)
    if err != nil {
        t.Fatal(err)
    }
    perm := stat.Mode().Perm()
    if perm != 0o600 {
        t.Errorf("want file mode 0o600, got 0%o", perm)
    }
    got, err := os.ReadFile(path)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Fatal(cmp.Diff(want, got))
    }
}

func TestWriteToFile_ReturnsErrorForUnwritableFile(t *testing.T){
    t.Parallel()
    path := "bogusdir/write_test.txt"
    err := writer.WriteToFile(path, []byte{})
    if err == nil {
        t.Fatal("want error when file not writable")
    }
}

func TestWriteToFile_ClobbersExistingFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/write_test.txt"
    err := os.WriteFile(path, []byte{4, 5, 6}, 0o600)
    if err != nil {
        t.Fatal(err)
    }

    want := []byte{1, 2, 3}
    err = writer.WriteToFile(path, want)
    if err != nil {
        t.Fatal(err)
    }
    got, err := os.ReadFile(path)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Fatal(cmp.Diff(want, got))
    }
}

func TestWriteToFile_ChangesPermsOnExistingFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/write_test.txt"
    err := os.WriteFile(path, []byte{4, 5, 6}, 0o645)
    if err != nil {
        t.Fatal(err)
    }

    want := []byte{1, 2, 3}
    err = writer.WriteToFile(path, want)
    if err != nil {
        t.Fatal(err)
    }
    stat, err := os.Stat(path)
    if err != nil {
        t.Fatal(err)
    }
    perm := stat.Mode().Perm()
    if perm != 0o600 {
        t.Errorf("want file mode 0o600, got 0%o", perm)
    }
}