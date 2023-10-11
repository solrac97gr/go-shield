package portscanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type PortScanner struct{}

func NewPortScanner() *PortScanner {
	return &PortScanner{}
}

func (p *PortScanner) Scan(host string, from, limit int) {
	existOpenPort := false
	if limit == 0 {
		limit = 65535
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 1000) // Adjust the number of concurrent goroutines based on your system's capabilities

	for i := from; i < limit; i++ {
		semaphore <- struct{}{} // Acquire a slot from the semaphore
		wg.Add(1)
		go func(port int) {
			defer func() {
				<-semaphore // Release the slot back to the semaphore
				wg.Done()
			}()

			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 2*time.Second) // Set a timeout value for connection attempts
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
