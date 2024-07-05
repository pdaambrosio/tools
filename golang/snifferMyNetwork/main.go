package main

import (
	"fmt"

	"github.com/pdaambrosio/tools/snifferMyNetwork/src"
)

// The main function prompts the user to enter a network interface name and BPF filter, then calls
// functions to search for network interfaces and capture packets on the specified interface.
func main() {
	var deviceName string
	var setFilter string
	fmt.Print("Enter the network interface name: ")
	fmt.Scanln(&deviceName)
	fmt.Print("Enter the BPF filter: ")
	fmt.Scanln(&setFilter)

	// Call the function `FindInterfaces` from the package `src` to search for a network interface
	fmt.Println("Searching for network interfaces...")
	src.FindInterfaces(deviceName)

	fmt.Println()

	// Call the function `LiveCapture` from the package `src` to capture packets on the network interface
	fmt.Println("Capturing packets...")
	src.LiveCapture(deviceName, setFilter)
}
