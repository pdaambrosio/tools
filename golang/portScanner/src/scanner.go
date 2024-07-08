package src

import (
	"fmt"
	"net"
	"sync"
)

// The ScanPorts function scans a range of ports on a given IP address to check for open ports.
func ScanPorts(ip string, portRange int) {
	var wg sync.WaitGroup
	fmt.Println("Scanning IP: ", ip)

	for i := 1; i <= portRange; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ip, j)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				// Port is closed
				fmt.Println("Port is closed or filtered: ", j)
				return
			}

			conn.Close()
			fmt.Printf("%d Open\n", j)
		}(i)
	}
	wg.Wait()
}
