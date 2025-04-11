package diskstats

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
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

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
	usedPercentage := (float64(used) / float64(total)) * 100

	return &DiskStats{
		TotalSpace:     total,
		FreeSpace:      free,
		UsedPercentage: usedPercentage,
	}, nil
	return nil, fmt.Errorf("disk stats not found for directory: %s", targetDir)
	return nil, fmt.Errorf("disk stats not found for directory: %s. Ensure the directory exists and you have sufficient permissions to access it", targetDir)
