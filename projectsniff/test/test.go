package main

import (
	"fmt"
	"projectsniff/readpacket"
	"strings"
)

func main() {
	devices, err := readpacket.Findalldevice()
	if err != nil {
		panic(err)
	}
	var interf string
	for _, device := range devices {
		if strings.Contains(device.Name, "BCBB1DA1") {
			interf = device.Name
			fmt.Println(interf)
			break
		}
	}

}
