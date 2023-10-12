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

// Ensure ConsoleClient implements the clients.Client interface.
var _ clients.Client = &ConsoleClient{}

// GetHost prompts the user to enter the hostname to check and returns the input.
func (c *ConsoleClient) GetHost() (host string) {
	return c.GetString("Enter the hostname to check: ")

}

// GetImagePath prompts the user to enter the image path to check and returns the input.
func (c *ConsoleClient) GetImagePath() (imagePath string) {
	return c.GetString("Enter the image path: ")

}

// GetFilePath prompts the user to enter the file path to check and returns the input.
func (c *ConsoleClient) GetFilePath() (filepath string) {
	return c.GetString("Enter the file path of the file: ")

}

// GetHash prompts the user to enter the hash to check and returns the input.
func (c *ConsoleClient) GetHash() (hash string) {
	return c.GetString("Enter the expected hash of the file: ")
}

// GetOpenPortsUserInfo prompts the user to enter the host, starting port, and limit for scanning open ports and returns the input.
func (c *ConsoleClient) GetOpenPortsUserInfo() (host string, from int, limit int) {
	host = c.GetString("Enter the host: ")
	from = c.GetInt("Enter the starting port: ")
	limit = c.GetInt("Enter the limit: ")
	return host, from, limit
}

func (c *ConsoleClient) GetPasswordGenerationInfo() (size int, withCapitalizedChar, withNumbers, withSpecialChar bool) {
	size = c.GetInt("Enter size of the password: ")
	withCapitalizedChar = c.GetBool("Use Capitalized letters? [default:Y] (Y or N): ")
	withNumbers = c.GetBool("Use Numbers? [default:Y] (Y or N): ")
	withSpecialChar = c.GetBool("Use Symbols? [default:Y] (Y or N):")
	return
}

// GetInt prompts the user to enter a number based on the provided message and returns the input.
func (c *ConsoleClient) GetInt(msg string) (number int) {
	fmt.Print(msg)
	fmt.Scanln(&number)
	return
}

// GetString prompts the user to enter a string based on the provided message and returns the input.
func (c *ConsoleClient) GetString(msg string) (str string) {
	fmt.Print(msg)
	fmt.Scanln(&str)
	return
}

// GetBool prompts the user to enter a boolean value based on the provided message and returns the input.
// The function displays the message to the user and waits for their input.
// The user's input is case-insensitive and can be "Y", "Yes", "yes" for true, or "N", "No", "no" for false.
// If the user's input does not match any of the expected values, the function returns true by default.
// If the provided message is an empty string, the function returns true.
func (c *ConsoleClient) GetBool(msg string) (is bool) {
	var answer string
	fmt.Print(msg)
	fmt.Scanln(&answer)
	if answer == "Y" || answer == "Yes" || answer == "yes" {
		return true
	}
	if answer == "N" || answer == "No" || answer == "no" {
		return false
	}
	if msg == "" {
		return true
	}
	return true
}

func (c *ConsoleClient) GetDecryptFileInfo() (filePath string, password string) {
	filePath = c.GetString("Enter the file path: ")
	password = c.GetString("Enter the file password: ")
	return
}
