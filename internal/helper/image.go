package helper

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func ResizeImage(img image.Image, targetWidth uint) (image.Image, error) {
	originalWidth := img.Bounds().Dx()

	if originalWidth < int(targetWidth) {
		return img, nil
	}

	return resize.Resize(targetWidth, 0, img, resize.Bicubic), nil
}

func EncodeImage(img image.Image, outputFile *os.File, format string) error {
	switch format {
	case "jpeg", "jpg":
		err := jpeg.Encode(outputFile, img, nil)
		if err != nil {
			return fmt.Errorf("failed to encode jpeg: %v", err)
		}
	case "png":
		err := png.Encode(outputFile, img)
		if err != nil {
			return fmt.Errorf("failed to encode png: %v", err)
		}
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
	return nil
}

func SaveResizedImage(path string, img image.Image, format string) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return EncodeImage(img, outputFile, format)
}
