package orchestrator

import (
	"fmt"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"

	"github.com/solrac97gr/go-shield/clients"
	fileencryption "github.com/solrac97gr/go-shield/tools/file-encryptation"
	fileintegritychecker "github.com/solrac97gr/go-shield/tools/file-integrity-checker"
	ipchecker "github.com/solrac97gr/go-shield/tools/ip-checker"
	"github.com/solrac97gr/go-shield/tools/metadata"
	passwordgenerator "github.com/solrac97gr/go-shield/tools/password-generator"
	portscanner "github.com/solrac97gr/go-shield/tools/port-scanner"
	whoischecker "github.com/solrac97gr/go-shield/tools/whois-checker"
)

type Orchestrator struct {
	Client               clients.Client
	IPChecker            *ipchecker.IPChecker
	PortScanner          *portscanner.PortScanner
	WhoIsChecker         *whoischecker.WhoIsChecker
	FileIntegrityChecker *fileintegritychecker.FileIntegrityChecker
	MetaDataCleanner     *metadata.MetaDataTools
	PasswordGenerator    *passwordgenerator.PasswordGenerator
	FileEncryptor        *fileencryption.FileEncryptor
}

func NewOrchestrator(c clients.Client, ipChecker *ipchecker.IPChecker, portScanner *portscanner.PortScanner, whoisCheck *whoischecker.WhoIsChecker, fileIntegrityChecker *fileintegritychecker.FileIntegrityChecker, metadataCleanner *metadata.MetaDataTools, passwordGenerator *passwordgenerator.PasswordGenerator, fileEncryptor *fileencryption.FileEncryptor) *Orchestrator {
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
	if metadataCleanner == nil {
		panic("nil metadata cleanner")
	}
	if passwordGenerator == nil {
		panic("nil password generator")
	}
	if fileEncryptor == nil {
		panic("nil file encryptor")
	}
	return &Orchestrator{
		Client:               c,
		IPChecker:            ipChecker,
		PortScanner:          portScanner,
		WhoIsChecker:         whoisCheck,
		FileIntegrityChecker: fileIntegrityChecker,
		MetaDataCleanner:     metadataCleanner,
		PasswordGenerator:    passwordGenerator,
		FileEncryptor:        fileEncryptor,
	}
}

func (o *Orchestrator) showMenu() {
	fmt.Println("Menu 📖:")
	fmt.Println("1. Check IP 🔍")
	fmt.Println("2. Scan Ports 📡")
	fmt.Println("3. Who is 🤔")
	fmt.Println("4. File integrity 🗒️")
	fmt.Println("5. Metadata Manipulation 🌃")
	fmt.Println("6. Generate password 🔒")
	fmt.Println("7. File encryptation 🛡️")
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
	case 5:
		o.showMetadataSubMenu()
	case 6:
		o.generatePassword()
	case 7:
		o.showFileEncryptationSubMenu()
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

func (o *Orchestrator) showMetadataSubMenu() {
	fmt.Println("Sub Menú 📖:")
	fmt.Println("1. Clean EXIF Metadata (IOS Images) 🌃")
	fmt.Println("2. Read EXIF Metadata (IOS Images) 🌃")
	fmt.Println("3. Return ⬅️")
	fmt.Println("0. Exit ❌")

	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		o.cleanEFIXMetadata()
	case 2:
		o.readEFIXMetadata()
	case 3:
		o.showMenu()
	default:
		fmt.Println("Invalid choice")
	}
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

func (o *Orchestrator) cleanEFIXMetadata() {
	imagepath := o.Client.GetImagePath()
	err := o.MetaDataCleanner.CleanEXIFMetaData(imagepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data cleanned succesfully 🧹")
}

func (o *Orchestrator) readEFIXMetadata() {
	imagepath := o.Client.GetImagePath()
	result, err := o.MetaDataCleanner.ReadEXIFMetaData(imagepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func (o *Orchestrator) generatePassword() {
	size, withCapitalizedChar, withNumbers, withSpecialChar := o.Client.GetPasswordGenerationInfo()
	password := o.PasswordGenerator.GenerateSafePassword(size, withCapitalizedChar, withNumbers, withSpecialChar)
	fmt.Printf("Password generated: %s\n", password)
}

func (o *Orchestrator) showFileEncryptationSubMenu() {
	fmt.Println("Sub Menú 📖:")
	fmt.Println("1. Encrypt File 🔒")
	fmt.Println("2. Decrypt File 🔑")
	fmt.Println("3. Return ⬅️")
	fmt.Println("0. Exit ❌")

	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		o.encryptFile()
	case 2:
		o.decryptFile()
	case 3:
		o.showMenu()
	default:
		fmt.Println("Invalid choice")
	}
}

func (o *Orchestrator) encryptFile() {
	filePath := o.Client.GetFilePath()
	outputPath, password, err := o.FileEncryptor.EncryptFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Output: %s", outputPath)
	fmt.Println()
	fmt.Printf("Key: %s", password)
	fmt.Println()
}

func (o *Orchestrator) decryptFile() {
	filePath, password := o.Client.GetDecryptFileInfo()
	outputPath, err := o.FileEncryptor.DecryptFile(filePath, password)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Output: %s", outputPath)
	fmt.Println()
}
