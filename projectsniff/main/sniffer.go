package main

import (
	"fmt"
	"projectsniff/readpacket"
	"strings"
)

func getinterface(pattern string) (interf string, err error) {
	devices, err := readpacket.Findalldevice()
	if err != nil {
		panic(err)
	}
	for _, device := range devices {
		if strings.Contains(device.Name, pattern) {
			interf = device.Name
			fmt.Println(interf)
		}
	}
	return
}

func main() {
	var x string
	device, err := getinterface("BCBB1DA1")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The interface being sniffer", device)
		go func() {
			readpacket.Sniffing(device)
		}()
		fmt.Scanf("%v", &x)
	}

}
