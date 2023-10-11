package metadata

import (
	"fmt"
	"image"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

type MetaDataTools struct{}

func NewMetaDataCleanner() *MetaDataTools {
	return &MetaDataTools{}
}

func (m *MetaDataTools) CleanEXIFMetaData(imagePath string) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Create a new image without the EXIF metadata
	newImg := imaging.Clone(img)

	// Save the new image back to the file
	err = imaging.Save(newImg, imagePath)
	if err != nil {
		return fmt.Errorf("failed to save modified image: %v", err)
	}

	return nil
}

func (m *MetaDataTools) ReadEXIFMetaData(imagePath string) (string, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the EXIF metadata

	exifData, err := exif.Decode(file)
	if err != nil {
		if err.Error() == "EOF" {
			return "No data", nil
		}
		if err.Error() == "exif: failed to find exif intro marker" {
			return "No data", nil
		}
		return "", fmt.Errorf("failed to decode EXIF metadata: %v", err)
	}

	return exifData.String(), nil
}
