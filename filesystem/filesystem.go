package filesystem

import (
	"os"
	"path/filepath"
	"sort"
)

type FilesystemEntry struct {
	Path string
	Size int64
}

type Filesystem struct {
	Entries []FilesystemEntry
}

func WalkDirectory(dir string) ([]FilesystemEntry, error) {
	entries := []FilesystemEntry{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			entries = append(entries, FilesystemEntry{Path: path, Size: info.Size()})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Size > entries[j].Size
	})
	return entries, nil
}
