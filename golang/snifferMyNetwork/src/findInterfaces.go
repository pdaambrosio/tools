package src

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

// The function `findInterfaces` searches for a specific network interface by name and prints its
// details if found.
func FindInterfaces(deviceName string) string {
	var Found bool = false

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("Unable to fecth network interfaces: ", err)
	}

	for _, ifDevice := range devices {
		if ifDevice.Name == deviceName {
			Found = true

			if len(ifDevice.Addresses) < 2 {
				log.Fatal("Network interface does not have IPv4 and IPv6 addresses or is not up")
			}

			var devName string = ifDevice.Name
			var devDesc string = ifDevice.Description
			var ipv4, ipv6 string = ifDevice.Addresses[0].IP.String(), ifDevice.Addresses[1].IP.String()

			fmt.Println("Network interface found:", devName)
			fmt.Println("Network interface description:", devDesc)
			fmt.Println("Network interface IPv4 addresses:", ipv4)
			fmt.Println("Network interface IPv6 addresses:", ipv6)
		}
	}

	if !Found {
		log.Fatal("Network interface not found")
	}

	return deviceName
}
