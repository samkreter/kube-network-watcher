package tcpdump

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
)

var (
	defaultDevice     = "eth0"

	// 65535 for full packet, 1024 for header
	snapshotLen int32  = 65535
	promiscuous    = false
	// Negative timeout always flushes packets
	timeout      = time.Second * -1
)


// StartPacketDump dumps all packets to the screen
// TODO: Allow passing customer BPF filters
func StartPacketDump() error {
	device, err := getDefaultNetworkInterface()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Using network interface: %s", device)

	handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Just look for tcp
	handle.SetBPFFilter("tcp")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	log.Info("Starting packet dump")

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	return nil
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
    	if strings.EqualFold(device.Name, defaultDevice){
			return defaultDevice, nil
		}
    }

    return devices[0].Name, nil
}


func getetstate() {
	out, err := exec.Command("netstat", "-tan").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}