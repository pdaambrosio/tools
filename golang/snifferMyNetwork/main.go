package main

import (
	"fmt"

	"github.com/pdaambrosio/tools/src"
)

func main() {
	// Call the function `FindInterfaces` from the package `src` to search for a network interface
	fmt.Println("Searching for network interfaces...")
	src.FindInterfaces("wlp8s0")
}
