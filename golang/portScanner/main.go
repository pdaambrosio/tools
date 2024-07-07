package main

import (
	"flag"

	"github.com/pdaambrosio/tools/portScanner/src"
)

// The main function in Go parses command line flags for IP address and port range, then scans the
// specified ports on the given IP address.
func main() {
	var ip string
	var portRange int

	flag.StringVar(&ip, "i", "", "The IP address to scan, sample 192.168.15.1 (required) ")
	flag.IntVar(&portRange, "p", 1024, "The range of ports to scan (required)")
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(ip) == 0 {
		flag.Usage()
		return
	}


	src.ScanPorts(ip, portRange)
}
