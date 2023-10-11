// Package whoischecker provides a simple whois lookup functionality by performing a TXT lookup on the whois subdomain of a given host.
package whoischecker

import (
	"github.com/likexian/whois"
)

// WhoIsChecker represents a whois checker object.
type WhoIsChecker struct{}

// NewWhoIsChecker creates a new WhoIsChecker instance.
func NewWhoIsChecker() *WhoIsChecker {
	return &WhoIsChecker{}
}

// WhoIs performs a whois lookup on the given host and prints the result.
func (w *WhoIsChecker) WhoIs(host string) (string, error) {
	result, err := whois.Whois(host)
	if err != nil {
		return "", err
	}
	return result, nil
}
