//go:build windows
// +build windows

package filesystem

import (
	"os"
	"path/filepath"
	"sort"
)

// walkDirectory traverses a directory and collects file entries (Windows-specific implementation).
func walkDirectory(dir string) ([]FilesystemEntry, error) {
	entries := []FilesystemEntry{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories and collect only files
		if !info.IsDir() {
			entries = append(entries, FilesystemEntry{
				Path: path,
				Size: info.Size(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort files by size in descending order
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Size > entries[j].Size
	})

	return entries, nil
}
