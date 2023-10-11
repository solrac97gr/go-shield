package main

import (
	"github.com/solrac97gr/go-shield/clients/console"
	"github.com/solrac97gr/go-shield/pkg/orchestrator"
	fileintegritychecker "github.com/solrac97gr/go-shield/tools/file-integrity-checker"
	ipchecker "github.com/solrac97gr/go-shield/tools/ip-checker"
	portscanner "github.com/solrac97gr/go-shield/tools/port-scanner"
	whoischecker "github.com/solrac97gr/go-shield/tools/whois-checker"
)

func main() {
	// Initialize the client
	console := console.NewConsoleClient()

	// Initialize the tools
	ipcheck := ipchecker.NewIPChecker()
	portscan := portscanner.NewPortScanner()
	whocheck := whoischecker.NewWhoIsChecker()
	ficheck := fileintegritychecker.NewFileIntegrityChecker()

	// Initialize orchestator
	orch := orchestrator.NewOrchestrator(
		console,
		ipcheck,
		portscan,
		whocheck,
		ficheck,
	)

	// Start
	orch.Start()
}
