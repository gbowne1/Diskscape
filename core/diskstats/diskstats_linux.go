//go:build !windows
// +build !windows

package diskstats

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// GetDiskStats retrieves disk usage statistics for the specified directory on POSIX systems.
func GetDiskStats(targetDir string) (*DiskStats, error) {
	var stat unix.Statfs_t
	err := unix.Statfs(targetDir, &stat)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve disk stats: %w", err)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	usedPercentage := (float64(used) / float64(total)) * 100

	return &DiskStats{
		TotalSpace:     total,
		FreeSpace:      free,
		UsedPercentage: usedPercentage,
	}, nil
}
