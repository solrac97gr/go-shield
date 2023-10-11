package fileintegritychecker

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

type FileIntegrityChecker struct{}

func NewFileIntegrityChecker() *FileIntegrityChecker {
	return &FileIntegrityChecker{}
}

func (fic *FileIntegrityChecker) CalculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString, nil
}

func (fic *FileIntegrityChecker) VerifyFileIntegrity(filePath string, expectedHash string) (bool, error) {
	actualHash, err := fic.CalculateFileHash(filePath)
	if err != nil {
		return false, err
	}

	return actualHash == expectedHash, nil
}
