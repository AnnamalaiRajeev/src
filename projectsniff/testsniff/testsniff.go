package main

import (
	"/projectsniff/flags"
	"fmt"
	"projectsniff/readpacket"
	"strings"
)


interfacepattern  := flags.Flag_{"filename","test.txt","The name of the input file"}


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
