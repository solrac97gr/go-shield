package fileencryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type FileEncryptor struct{}

func NewFileEncryptor() *FileEncryptor {
	return &FileEncryptor{}
}

func (fe *FileEncryptor) EncryptFile(filePath string) (outputPath, generatedPassword string, err error) {
	outputPath = getEncryptedFilePath(filePath)
	generatedPassword = generateRandomPassword()

	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}

	key := []byte(generatedPassword)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	encrypted := make([]byte, aes.BlockSize+len(input))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], input)

	err = ioutil.WriteFile(outputPath, encrypted, 0644)
	if err != nil {
		return "", "", err
	}

	return outputPath, generatedPassword, nil
}

func (fe *FileEncryptor) DecryptFile(filePath string, password string) (outputPath string, err error) {
	outputPath = getDecryptedFilePath(filePath)

	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	key := []byte(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(input) < aes.BlockSize {
		return "", errors.New("invalid file")
	}

	iv := input[:aes.BlockSize]
	encrypted := input[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	err = ioutil.WriteFile(outputPath, encrypted, 0644)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

func getEncryptedFilePath(filePath string) string {
	ext := filepath.Ext(filePath)
	fileName := strings.TrimSuffix(filepath.Base(filePath), ext)
	return filepath.Join(filepath.Dir(filePath), fileName+".encrypted")
}

func getDecryptedFilePath(filePath string) string {
	ext := filepath.Ext(filePath)
	fileName := strings.TrimSuffix(filepath.Base(filePath), ext)
	return filepath.Join(filepath.Dir(filePath), fileName+".decrypted")
}

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
