// https://tour.golang.org/methods/18
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (byte IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", byte[0], byte[1], byte[2], byte[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
