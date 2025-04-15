package filesystem

// FilesystemEntry represents a single file or directory with its path and size.
type FilesystemEntry struct {
	Path string // Full path to the file or directory
	Size int64  // Size in bytes (0 for directories if not calculated)
}

// WalkDirectory is an interface for traversing directories.
// Each platform-specific implementation provides its own version of this function.
func WalkDirectory(dir string) ([]FilesystemEntry, error) {
	return walkDirectory(dir) // Delegates to platform-specific implementations
}
