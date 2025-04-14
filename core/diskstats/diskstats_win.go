//go:build windows
// +build windows

package diskstats

import (
	"fmt"
	"syscall"
	"unsafe"
)

// GetDiskStats retrieves disk usage statistics for the specified directory on Windows systems.
func GetDiskStats(targetDir string) (*DiskStats, error) {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	getDiskFreeSpaceEx := kernel32.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64

	_, _, err := getDiskFreeSpaceEx.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(targetDir))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
	)
	if err != syscall.Errno(0) {
		return nil, fmt.Errorf("failed to retrieve disk stats: %w", err)
	}

	usedBytes := totalNumberOfBytes - totalNumberOfFreeBytes
	usedPercentage := (float64(usedBytes) / float64(totalNumberOfBytes)) * 100

	return &DiskStats{
		TotalSpace:     totalNumberOfBytes,
		FreeSpace:      totalNumberOfFreeBytes,
		UsedPercentage: usedPercentage,
	}, nil
}
