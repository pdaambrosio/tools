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
			fmt.Println("Network interface found: ", ifDevice.Name)
			fmt.Println("Network interface description: ", ifDevice.Description)
			fmt.Println("Network interface addresses: ", ifDevice.Addresses)
		}
	}

	if !Found {
		log.Panicln("Network interface not found")
	}

	return deviceName
}
