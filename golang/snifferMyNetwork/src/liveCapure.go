package src

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// The `LiveCapture` function opens a network interface for live packet capture, sets a BPF filter, and
// prints captured packets until the capture is finished.
func LiveCapture(deviceName string, setFilter string) string {
	handler, err := pcap.OpenLive(deviceName, 1600, false, pcap.BlockForever)
	if err != nil {
		log.Fatalln("Unable to open network interface for live capture: ", err)
	}

	defer handler.Close()

	if err := handler.SetBPFFilter(setFilter); err != nil {
		log.Fatalln("Unable to set BPF filter: ", err)
	}

	packetSource := gopacket.NewPacketSource(handler, handler.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	return "Live capture finished"
}
