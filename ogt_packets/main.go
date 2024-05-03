package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln("Unable to fetch network interfaces")
	}

	for _, device := range devices {
		if len(device.Addresses) != 0 {
			fmt.Printf("XXXXXXXXXXXXX\n")
			fmt.Printf("%s\n", device.Name)
			fmt.Printf("%s\n", device.Description)
			fmt.Printf("%s\n", device.Addresses)
		}
	}

	handle, err := pcap.OpenLive(
		"\\Device\\NPF_{C1691B05-0ECF-4620-994E-4BC5B9EC289D}",
		1600,
		false,
		pcap.BlockForever)

	if err != nil {
		log.Panicln("Unable to fetch network interfaces")
	}

	defer handle.Close()

	if err := handle.SetBPFFilter("tcp and port 443"); err != nil {
		log.Panic(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		fmt.Println(packet)
	}

}
