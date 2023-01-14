package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run sniffer.go [interface]")
		os.Exit(1)
	}

	// Open interface for capturing packets
	handle, err := pcap.OpenLive(os.Args[1], 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
