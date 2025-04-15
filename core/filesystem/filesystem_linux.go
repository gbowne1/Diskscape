//go:build linux || unix || darwin
// +build linux unix darwin

package filesystem

import (
	"os"
	"path/filepath"
	"sort"
)

// walkDirectory traverses a directory and collects file entries (POSIX/Linux/Unix-specific implementation).
func walkDirectory(dir string) ([]FilesystemEntry, error) {
	entries := []FilesystemEntry{}

	// Use filepath.WalkDir introduced in Go 1.16 for improved traversal
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Return error with context if traversal fails
			return err
		}
		// Skip directories and collect only file entries
		if !d.IsDir() {
			info, statErr := d.Info() // Retrieve os.FileInfo for size
			if statErr != nil {
				return statErr
			}
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
