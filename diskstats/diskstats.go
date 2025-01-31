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

func GetDiskStats(targetDir string) (*DiskStats, error) {
	cmd := exec.Command("df", "-k", targetDir)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute df command: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, targetDir) {
			parts := strings.Fields(line)
			total, _ := strconv.ParseUint(parts[1], 10, 64)
			free, _ := strconv.ParseUint(parts[2], 10, 64)
			usedPercentage, _ := strconv.ParseFloat(parts[4], 32)

			return &DiskStats{
				TotalSpace:     total,
				FreeSpace:      free,
				UsedPercentage: usedPercentage,
			}, nil
		}
	}

	return nil, fmt.Errorf("disk stats not found for directory: %s", targetDir)
}
