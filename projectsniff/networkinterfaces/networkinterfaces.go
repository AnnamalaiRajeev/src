package networkinterfaces

import (
	"net"
)

// InterfacesNameList..
type InterfacesNameList struct {
	name []string
}

// LocalAddresses Returns a pointer to the list data structure containing interface names
func LocalAddresses() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var k int
	x := InterfacesNameList{}
	x.name = make([]string, 10, 10)
	for j, i := range ifaces {
		x.name[j] = i.Name
		k = j
	}
	return x.name[:k+1], err
}
