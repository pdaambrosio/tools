package src

import (
	"sort"
)

// The ScanPorts function scans a range of ports on a given IP address and returns a list of open
// ports.
func ScanPorts(ip string, portRange int) []int {
	// Create a channel to hold the port numbers
	portChan := make(chan int, portRange)
	results := make(chan int, portRange)
	var openPorts []int

	// Start the worker goroutines
	for i := 0; i < 100; i++ {
		go Worker(portChan, ip, results)
	}

	// Send the port numbers to the worker goroutines
	for i := 1; i <= portRange; i++ {
		portChan <- i
	}

	// Close the portChan channel
	close(portChan)

	// Collect the results
	for i := 1; i <= portRange; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// Sort the openPorts slice
	sort.Ints(openPorts)

	return openPorts
}
