// Package portscanner provides a simple port scanner that scans for open TCP ports on a given host.
package portscanner

import (
	"fmt"
	"net"
	"sync"
)

// PortScanner represents a port scanning object.
type PortScanner struct{}

// NewPortScanner creates a new PortScanner instance.
func NewPortScanner() *PortScanner {
	return &PortScanner{}
}

// scan performs the port scanning on a given host.
func (p *PortScanner) Scan(host string, from, limit int) {
	existOpenPort := false

	if limit == 0 {
		limit = 65535
	}

	var wg sync.WaitGroup
	for i := from; i < limit; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return
			}
			conn.Close()
			existOpenPort = true
			fmt.Println("Port", port, "is open")
		}(i)
	}
	wg.Wait()

	if !existOpenPort {
		fmt.Printf("No ports open for %s", host)
	}
}
