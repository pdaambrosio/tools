package src

import (
	"fmt"
	"net"
)

// The ScanPorts function scans a range of ports on a given IP address to check for open ports.
func ScanPorts(ip string, portRange int) {
	fmt.Println("Scanning IP: ", ip)

	for i := 1; i <= portRange; i++ {
		address := fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			// Port is closed
			fmt.Println("Port closed: ", i)
			continue
		}

		conn.Close()
		fmt.Printf("%d Open\n", i)
	}
}
