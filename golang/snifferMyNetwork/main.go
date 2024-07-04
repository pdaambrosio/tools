package main

import (
	"fmt"

	"github.com/pdaambrosio/tools/src"
)

func main() {
	var deviceName string
	fmt.Print("Enter the network interface name: ")
	fmt.Scanln(&deviceName)

	// Call the function `FindInterfaces` from the package `src` to search for a network interface
	fmt.Println("Searching for network interfaces...")
	src.FindInterfaces(deviceName)

	fmt.Println()

	// Call the function `LiveCapture` from the package `src` to capture packets on the network interface
	fmt.Println("Capturing packets...")
	src.LiveCapture(deviceName, "tcp and port 80")
}
