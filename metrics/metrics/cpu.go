// metrics/cpu.go
package metrics

import (
	"github.com/shirou/gopsutil/cpu"
	"fmt"
)

// CPUUsage represents the CPU metrics response
type CPUUsage struct {
	CPUUsage float64 `json:"cpu_usage"`
}

// GetCPUUsage fetches the CPU usage using the gopsutil/cpu package
func GetCPUUsage() (*CPUUsage, error) {
	// Fetch the CPU usage percentage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, fmt.Errorf("Error fetching CPU usage: %v", err)
	}

	// Return the CPU usage
	return &CPUUsage{
		CPUUsage: cpuPercent[0],
	}, nil
}
