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

// GetHost prompts the user to enter the hostname to check and returns the input.
func (*ConsoleClient) GetHost() (host string) {
	fmt.Print("Enter the hostname to check: ")
	fmt.Scanln(&host)
	return
}

// GetFilePath prompts the user to enter the file path to check and returns the input.
func (*ConsoleClient) GetFilePath() (filepath string) {
	fmt.Print("Enter the file path of the file: ")
	fmt.Scanln(&filepath)
	return
}

// GetHash prompts the user to enter the hash to check and returns the input.
func (*ConsoleClient) GetHash() (hash string) {
	fmt.Print("Enter the expected hash of the file: ")
	fmt.Scanln(&hash)
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
