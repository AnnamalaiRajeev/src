package readpacket

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var icmpfilter string = "src host 0.0.0.0 and icmp"
var dnsfilter string = "udp and port 53"
var (
	snaplen int32 = 65535
	promisc bool  = false
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

// "Sniffing on a interface"
func Sniffing(interfacename string) {

	handle, err = pcap.OpenLive(interfacename, snaplen, promisc, timeout)
	if err != nil {
		panic(err)
	} else {
		defer handle.Close()
		err = handle.SetBPFFilter(dnsfilter)
		if err != nil {
			log.Fatal(err)
		}
	}
	packetsource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetsource.Packets() {
		fmt.Println(packet)

	}
}

// Finds all interfaces attached to the host
func Findalldevice() ([]pcap.Interface, error) {
	var devices []pcap.Interface
	devices, err = pcap.FindAllDevs()
	if err != nil {
		panic(err)
	} else {
		return devices, err
	}

}
