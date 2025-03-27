package collector

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// After structuring we are able to Populate.
type Payload struct {
	Disk   DiskMetrics   `json:"disk"`
	Memory MemoryMetrics `json:"memory"`
	OS     OSMetrics     `json:"OS"`
	CPU    CPUUsage      `"json:CPU"`
}

type DiskMetrics struct {
	Total          uint64 `json:"total"`
	Used           uint64 `json:"used"`
	IopsInProgress uint64 `json:"iopsInProgress"`
}

type MemoryMetrics struct {
	SwapTotal    uint64 `json:"swap_total"`
	SwapUsed     uint64 `json:"swap_used"`
	VirtualTotal uint64 `json:"virtual_total"`
	VirtualUsed  uint64 `json:"virtual_used"`
	Buffers      uint64 `json:"buffers"`
	Cached       uint64 `json:"cached"`
}

type OSMetrics struct {
	Uptime          uint64 `json:"uptime"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

type CPUUsage struct {
	CPUUsage float64 `json:"cpu_usage"`
}

// Define interfaces for dependency injection
type DiskMetricsProvider interface {
	GetDiskMetrics() (DiskMetrics, error)
}

type MemoryMetricsProvider interface {
	GetMemoryMetrics() (MemoryMetrics, error)
}

type OSMetricsProvider interface {
	GetOSMetrics() (OSMetrics, error)
}

type CPUMetricsProvider interface {
	GetCPUUsage() (CPUUsage, error)
}

// Implement concrete system metric collection
type SystemMetrics struct{}

func (s SystemMetrics) GetDiskMetrics() (DiskMetrics, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return DiskMetrics{}, err
	}
	return DiskMetrics{Total: diskUsage.Total, Used: diskUsage.Used}, nil
}

func (s SystemMetrics) GetMemoryMetrics() (MemoryMetrics, error) {
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return MemoryMetrics{}, err
	}
	return MemoryMetrics{VirtualTotal: memoryStats.Total, VirtualUsed: memoryStats.Used}, nil
}

func (s SystemMetrics) GetOSMetrics() (OSMetrics, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return OSMetrics{}, err
	}
	return OSMetrics{Uptime: hostInfo.Uptime, Platform: hostInfo.Platform}, nil
}

func (s SystemMetrics) GetCPUUsage() (CPUUsage, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return CPUUsage{}, err
	}
	return CPUUsage{CPUUsage: cpuPercent[0]}, nil
}

// Inject dependencies
type MetricCollector struct {
	Disk   DiskMetricsProvider
	Memory MemoryMetricsProvider
	OS     OSMetricsProvider
	CPU    CPUMetricsProvider
}

func (mc *MetricCollector) Collect() (Payload, error) {
	diskMetrics, err := mc.Disk.GetDiskMetrics()
	if err != nil {
		return Payload{}, err
	}

	memoryMetrics, err := mc.Memory.GetMemoryMetrics()
	if err != nil {
		return Payload{}, err
	}

	osMetrics, err := mc.OS.GetOSMetrics()
	if err != nil {
		return Payload{}, err
	}

	cpuUsage, err := mc.CPU.GetCPUUsage()
	if err != nil {
		return Payload{}, err
	}

	return Payload{Disk: diskMetrics, Memory: memoryMetrics, OS: osMetrics, CPU: cpuUsage}, nil
}

// func main() {
// 	// Inject dependencies
// 	sysMetrics := SystemMetrics{}
// 	collector := MetricCollector{Disk: sysMetrics, Memory: sysMetrics, OS: sysMetrics, CPU: sysMetrics}

// 	// Collect metrics
// 	metrics, err := collector.Collect()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Collected Metrics:", metrics)
// }
