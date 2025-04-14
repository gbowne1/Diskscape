package diskstats

import (
	"testing"
)

func TestGetDiskStats(t *testing.T) {
	stats, err := GetDiskStats("/")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if stats.TotalSpace == 0 {
		t.Errorf("Expected non-zero TotalSpace, got %d", stats.TotalSpace)
	}
}

func GetDiskStats(s string) (any, any) {
	panic("unimplemented")
}

func TotalSpace() int {
	return 0
}
func FreeSpace() int {
	return 0
}
func UsedPercentage() int {
	return 0
}
