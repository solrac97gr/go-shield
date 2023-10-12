package fileencryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileEncryptor struct{}

func NewFileEncryptor() *FileEncryptor {
	return &FileEncryptor{}
}

func (fe *FileEncryptor) EncryptFile(filePath string) (outputPath string, password string, err error) {
	// Generate a random password
	password = generateRandomPassword()

	// Open the input file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer inputFile.Close()

	// Create the output file
	outputPath = getOutputFilePath(filePath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", "", err
	}
	defer outputFile.Close()

	// Generate a random key and IV
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "", "", err
	}
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return "", "", err
	}

	// Write the key and IV to the output file
	_, err = outputFile.Write(key)
	if err != nil {
		return "", "", err
	}
	_, err = outputFile.Write(iv)
	if err != nil {
		return "", "", err
	}

	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// Create the stream cipher with the IV
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt and write the file content to the output file
	buffer := make([]byte, 4096)
	for {
		n, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", "", err
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = outputFile.Write(buffer[:n])
		if err != nil {
			return "", "", err
		}
	}

	return outputPath, password, nil
}

func (fe *FileEncryptor) DecryptFile(filePath string, password string) (outputPath string, err error) {
	// Open the input file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	// Read the key and IV from the input file
	key := make([]byte, 32)
	_, err = inputFile.Read(key)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	_, err = inputFile.Read(iv)
	if err != nil {
		return "", err
	}

	// Create the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create the stream cipher with the IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create the output file
	outputPath = getDecryptedFilePath(filePath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Decrypt and write the file content to the output file
	buffer := make([]byte, 4096)
	for {
		n, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = outputFile.Write(buffer[:n])
		if err != nil {
			return "", err
		}
	}

	return outputPath, nil
}

// Helper function to generate a random password
func generateRandomPassword() string {
	const passwordLength = 16
	const passwordCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, passwordLength)
	_, err := rand.Read(password)
	if err != nil {
		panic(err)
	}

	for i := 0; i < passwordLength; i++ {
		password[i] = passwordCharset[int(password[i])%len(passwordCharset)]
	}

	return string(password)
}

// Helper function to get the decrypted file path
func getDecryptedFilePath(filePath string) string {
	ext := filepath.Ext(filePath)
	fileName := strings.TrimSuffix(filepath.Base(filePath), ext)
	return filepath.Join(filepath.Dir(filePath), fileName+".decrypted")
}

// Helper function to get the output file path
func getOutputFilePath(filePath string) string {
	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	return filepath.Join(filepath.Dir(filePath), fileName+".encrypt")
}
