package main

import (
	"github.com/solrac97gr/go-shield/clients/console"
	"github.com/solrac97gr/go-shield/pkg/orchestrator"
	ipchecker "github.com/solrac97gr/go-shield/tools/ip-checker"
	portscanner "github.com/solrac97gr/go-shield/tools/port-scanner"
)

func main() {
	// Initialize the client
	console := console.NewConsoleClient()

	// Initialize the tools
	ipcheck := ipchecker.NewIPChecker()
	portscan := portscanner.NewPortScanner()

	// Initialize orchestator
	orch := orchestrator.NewOrchestrator(console, ipcheck, portscan)

	// Start
	orch.Start()
}
