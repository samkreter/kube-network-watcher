package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"time"
)

var (
	defaultdDevice     = "eth0"
	snapshot_len int32  = 1024
	promiscuous    = false
	err          error
	timeout      = 30 * time.Second
	handle       *pcap.Handle
)

func StartPacketDump() {
	device, err := getDefaultNetworkInterface()
	if err != nil {
		log.Fatal(err)
	}

	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}


// getDefaultNetworkInterface use the default network device if present, otherwise use the first in the list
func getDefaultNetworkInterface() (string, error) {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        return "", err
    }

    if devices == nil || len(devices) < 1 {
    	return "", fmt.Errorf("no network devices found")
	}

    for _, device := range devices {
    	if strings.EqualFold(device.Name, defaultdDevice){
			return defaultdDevice, nil
		}
    }

    return devices[0].Name, nil
}