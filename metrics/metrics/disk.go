// metrics/disk.go
package metrics

import (
	"github.com/shirou/gopsutil/disk"
	"fmt"
)

// DiskMetrics represents the disk metrics response
type DiskMetrics struct {
	TotalDisk   uint64  `json:"total_disk"`
	FreeDisk    uint64  `json:"free_disk"`
	UsedDisk    uint64  `json:"used_disk"`
	DiskPercent float64 `json:"disk_percent"`
}

// GetDiskUsage fetches the disk usage using the gopsutil/disk package
func GetDiskUsage() (*DiskMetrics, error) {
	// Fetch the disk usage for the root directory
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("Error fetching Disk usage: %v", err)
	}

	// Return the disk metrics
	return &DiskMetrics{
		TotalDisk:   diskStat.Total,
		FreeDisk:    diskStat.Free,
		UsedDisk:    diskStat.Used,
		DiskPercent: diskStat.UsedPercent,
	}, nil
}
