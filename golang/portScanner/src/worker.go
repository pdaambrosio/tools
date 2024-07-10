package src

import (
	"fmt"
	"net"
)

// The Worker function concurrently checks for open ports within a specified range on a given address
// and sends the results back through a channel.
func Worker(portRange chan int, address string, results chan int) {
	for port := range portRange {
		addrPort := fmt.Sprintf("%s:%d", address, port)
		conn, err := net.Dial("tcp", addrPort)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- port
	}
}
