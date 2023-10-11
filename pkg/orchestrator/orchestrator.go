package orchestrator

import (
	"fmt"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"

	"github.com/solrac97gr/go-shield/clients"
	fileintegritychecker "github.com/solrac97gr/go-shield/tools/file-integrity-checker"
	ipchecker "github.com/solrac97gr/go-shield/tools/ip-checker"
	portscanner "github.com/solrac97gr/go-shield/tools/port-scanner"
	whoischecker "github.com/solrac97gr/go-shield/tools/whois-checker"
)

type Orchestrator struct {
	Client               clients.Client
	IPChecker            *ipchecker.IPChecker
	PortScanner          *portscanner.PortScanner
	WhoIsChecker         *whoischecker.WhoIsChecker
	FileIntegrityChecker *fileintegritychecker.FileIntegrityChecker
}

func NewOrchestrator(c clients.Client, ipChecker *ipchecker.IPChecker, portScanner *portscanner.PortScanner, whoisCheck *whoischecker.WhoIsChecker, fileIntegrityChecker *fileintegritychecker.FileIntegrityChecker) *Orchestrator {
	if c == nil {
		panic("nil client")
	}
	if ipChecker == nil {
		panic("nil ip checker")
	}
	if portScanner == nil {
		panic("nil port scanner")
	}
	if whoisCheck == nil {
		panic("nil whois checker")
	}
	if fileIntegrityChecker == nil {
		panic("nil file integrity checker")
	}
	return &Orchestrator{
		Client:               c,
		IPChecker:            ipChecker,
		PortScanner:          portScanner,
		WhoIsChecker:         whoisCheck,
		FileIntegrityChecker: fileIntegrityChecker,
	}
}

func (o *Orchestrator) showMenu() {
	fmt.Println("Menu 📖:")
	fmt.Println("1. Check IP 🔍")
	fmt.Println("2. Scan Ports 📡")
	fmt.Println("3. Who is 🤔")
	fmt.Println("4. File integrity 🗒️")
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
	case 3:
		o.whoIs()
	case 4:
		o.showFileIntegritySubMenu()
	default:
		fmt.Println("Invalid choice")
	}
}

func (o *Orchestrator) Start() {
	figure := figure.NewColorFigure("Go-Shield", "epic", "green", true)

	figure.Print()
	fmt.Println()

	o.showMenu()
}

func (o *Orchestrator) checkIP() {
	host := o.Client.GetHost()
	o.IPChecker.CheckIP(host)
}

func (o *Orchestrator) scanPorts() {
	host, from, limit := o.Client.GetOpenPortsUserInfo()
	o.PortScanner.Scan(host, from, limit)
}

func (o *Orchestrator) whoIs() {
	host := o.Client.GetHost()
	result, err := o.WhoIsChecker.WhoIs(host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func (o *Orchestrator) showFileIntegritySubMenu() {
	fmt.Println("Sub Menú 📖:")
	fmt.Println("1. Calculate Hash File 📝")
	fmt.Println("2. Verify File 🔍")
	fmt.Println("3. Return ⬅️")
	fmt.Println("0. Exit ❌")

	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		o.calculateFileHash()
	case 2:
		o.verifyFileHash()
	case 3:
		o.showMenu()
	default:
		fmt.Println("Invalid choice")
	}
}

func (o *Orchestrator) calculateFileHash() {
	fp := o.Client.GetFilePath()
	res, err := o.FileIntegrityChecker.CalculateFileHash(fp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hash: %s", res)
	fmt.Println()
}

func (o *Orchestrator) verifyFileHash() {
	fp := o.Client.GetFilePath()
	expectedHash := o.Client.GetHash()
	res, err := o.FileIntegrityChecker.VerifyFileIntegrity(fp, expectedHash)
	if err != nil {
		log.Fatal(err)
	}
	if !res {
		fmt.Println("The file has been compromised 🔴")
		return
	}
	fmt.Println("the file is secure 👍")
}
