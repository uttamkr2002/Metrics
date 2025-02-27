package client

import (
	"fmt"
	"log"
	"metrics/db"
	"metrics/metrics"
	"time"
)

// Initialize DB connection
func init() {
	db.InitDB() // Initialize the database connection
}

// Insert metrics into PostgreSQL
func insertMetrics(cpuUsage *metrics.CPUUsage, diskMetrics *metrics.DiskMetrics) error {
	fmt.Println("Hi, We are Inserting the data")
	query := `
        INSERT INTO metrics_data (cpu_usage, disk_percent, used_disk, free_disk, total_disk)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := db.GetDB().Exec(query, cpuUsage.CPUUsage, diskMetrics.DiskPercent, diskMetrics.UsedDisk, diskMetrics.FreeDisk, diskMetrics.TotalDisk)
	if err != nil {
		return fmt.Errorf("error inserting metrics into database: %w", err)
	}
	return nil
}

// fetch data

func periodicFetchMetrics() {
	fmt.Println("Hi, We are in Periodic fetch")
	ticker := time.NewTicker(60 * time.Second) // Run every 60 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			//Fetch CPU usage
			cpuUsage, err := metrics.GetCPUUsage() // calling client function
			if err != nil {
				log.Println("Error fetching CPU usage:", err)
				continue
			}
			fmt.Println(cpuUsage)

			// Fetch Disk usage
			diskMetrics, err := metrics.GetDiskUsage() // calling client
			if err != nil {
				log.Println("Error fetching Disk usage:", err)
				continue
			}
			fmt.Println(diskMetrics)
			// Store the fetched metrics in PostgreSQL
			if err := insertMetrics(cpuUsage, diskMetrics); err != nil {
				log.Println("Error storing metrics in DB:", err)
			}

			//time.Sleep(1 * time.Minute) // it will create problem, as it will pause the function for 60 sec, no calculation is  there, So Try Something better
		}

	}
}

// Add a Post request to send the data to server ** later on

func StoreRun() {
	fmt.Println("We are In Client !..")
	// Start the periodic fetching and storing in the background
	go periodicFetchMetrics()
	time.Sleep(10 * time.Minute) // ? issue
}
