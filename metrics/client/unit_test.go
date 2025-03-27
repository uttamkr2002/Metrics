package tests

import (
	"testing"

	Temp "metrics/Collector"
)

// we are testing here collect function

// Mock implementation for testing
type MockMetrics struct{}

func (m MockMetrics) GetDiskMetrics() (Temp.DiskMetrics, error) {
	return Temp.DiskMetrics{Total: 1000, Used: 500}, nil
}

func (m MockMetrics) GetMemoryMetrics() (Temp.MemoryMetrics, error) {
	return Temp.MemoryMetrics{VirtualTotal: 8000, VirtualUsed: 4000}, nil
}

func (m MockMetrics) GetOSMetrics() (Temp.OSMetrics, error) {
	return Temp.OSMetrics{Uptime: 10000, Platform: "Linux"}, nil
}

func (m MockMetrics) GetCPUUsage() (Temp.CPUUsage, error) {
	return Temp.CPUUsage{CPUUsage: 25.5}, nil
}

func TestMetricCollector_Collect(t *testing.T) {
	// Inject the mock implementation
	mockMetrics := MockMetrics{}
	collector := Temp.MetricCollector{
		Disk:   mockMetrics,
		Memory: mockMetrics,
		OS:     mockMetrics,
		CPU:    mockMetrics,
	}

	// Call the method we are testing
	metrics, err := collector.Collect()

	// Check for unexpected errors
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Validate collected metrics
	if metrics.Disk.Total != 1000 {
		t.Errorf("Expected Disk.Total = 1000, got %d", metrics.Disk.Total)
	}

	if metrics.Memory.VirtualTotal != 8000 {
		t.Errorf("Expected Memory.VirtualTotal = 8000, got %d", metrics.Memory.VirtualTotal)
	}

	if metrics.OS.Platform != "Linux" {
		t.Errorf("Expected OS.Platform = 'Linux', got %s", metrics.OS.Platform)
	}

	if metrics.CPU.CPUUsage != 25.5 {
		t.Errorf("Expected CPU.CPUUsage = 25.5, got %f", metrics.CPU.CPUUsage)
	}
}
