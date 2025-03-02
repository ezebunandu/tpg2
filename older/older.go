package older

import (
	"fmt"
	"io/fs"
	"os"
	"time"
)

func OlderFiles(fsys fs.FS, age time.Duration) (paths []string) {
	threshold := time.Now().Add(-age)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil || info.IsDir() {
			return nil
		}
		if info.ModTime().Before(threshold) {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}

const Usage = `Usage: older DURATION

Lists all files older than DURATION in the tree rooted at the current directory.

Example: older 24h
(lists all files last modified more than 24 hours ago)`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 1
	}

	age, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fsys := os.DirFS(".")
	paths := OlderFiles(fsys, age)
	for _, p := range paths {
		fmt.Println(p)
	}
	return 0
}
