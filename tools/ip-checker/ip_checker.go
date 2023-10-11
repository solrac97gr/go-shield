// Package ipchecker provides a simple IP checking functionality by resolving the IP address of a given host.
package ipchecker

import (
	"fmt"
	"net"
)

// IPChecker represents an IP checking object.
type IPChecker struct{}

// NewIPChecker creates a new IPChecker instance.
func NewIPChecker() *IPChecker {
	return &IPChecker{}
}

// CheckIP resolves the IP address of the given host and prints it.
func (i *IPChecker) CheckIP(host string) {
	ip, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The IP for %s is: %s\n", host, ip[0].String())
}
