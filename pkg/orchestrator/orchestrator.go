package orchestrator

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"

	"github.com/solrac97gr/go-shield/clients"
	ipchecker "github.com/solrac97gr/go-shield/tools/ip-checker"
	portscanner "github.com/solrac97gr/go-shield/tools/port-scanner"
)

type Orchestrator struct {
	Client      clients.Client
	IPChecker   *ipchecker.IPChecker
	PortScanner *portscanner.PortScanner
}

func NewOrchestrator(c clients.Client, ipChecker *ipchecker.IPChecker, portScanner *portscanner.PortScanner) *Orchestrator {
	if c == nil {
		panic("nil client")
	}
	if ipChecker == nil {
		panic("nil ip checker")
	}
	if portScanner == nil {
		panic("nil port scanner")
	}
	return &Orchestrator{
		Client:      c,
		IPChecker:   ipChecker,
		PortScanner: portScanner,
	}
}

func (o *Orchestrator) Start() {
	figure := figure.NewColorFigure("Go-Shield", "epic", "green", true)

	figure.Print()
	fmt.Println()

	fmt.Println("Menu 📖:")
	fmt.Println("1. Check IP 🔍")
	fmt.Println("2. Scan Ports 📡")
	fmt.Println("0. Exit ❌")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		o.checkIP()
	case 2:
		o.scanPorts()
	default:
		fmt.Println("Invalid choice")
	}
}

func (o *Orchestrator) checkIP() {
	host := o.Client.GetHostForCheckFromUser()
	o.IPChecker.CheckIP(host)
}

func (o *Orchestrator) scanPorts() {
	host, from, limit := o.Client.GetOpenPortsUserInfo()
	o.PortScanner.Scan(host, from, limit)
}
