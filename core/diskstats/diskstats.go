package diskstats

type DiskStats struct {
	TotalSpace     uint64
	FreeSpace      uint64
	UsedPercentage float64
}
