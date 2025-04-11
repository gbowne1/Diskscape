package diskstats

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type DiskStats struct {
	TotalSpace     uint64
	FreeSpace      uint64
	UsedPercentage float64
}

// GetDiskStats retrieves disk usage statistics for the specified directory.
// The targetDir parameter should be a valid directory path on the system.
// This function assumes the use of the 'df' command to retrieve disk statistics.
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
