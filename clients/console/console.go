// Package console provides a console client implementation for interacting with the shield library.
package console

import (
	"fmt"

	"github.com/solrac97gr/go-shield/clients"
)

// ConsoleClient represents a console client.
type ConsoleClient struct{}

// NewConsoleClient creates a new ConsoleClient instance.
func NewConsoleClient() *ConsoleClient {
	return &ConsoleClient{}
}

var _ clients.Client = &ConsoleClient{}

// GetHostForCheckFromUser prompts the user to enter the hostname to check and returns the input.
func (*ConsoleClient) GetHostForCheckFromUser() (host string) {
	fmt.Print("Enter the hostname to check: ")
	fmt.Scanln(&host)
	return
}

// GetOpenPortsUserInfo prompts the user to enter the host, starting port, and limit for scanning open ports and returns the input.
func (*ConsoleClient) GetOpenPortsUserInfo() (host string, from int, limit int) {
	fmt.Print("Enter the host: ")
	fmt.Scanln(&host)
	fmt.Print("Enter the starting port: ")
	fmt.Scanln(&from)
	fmt.Print("Enter the limit: ")
	fmt.Scanln(&limit)
	return
}
