package imageTool

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

type Base64Image struct {
	Data image.Image
	Size int
	Ext  string
	Name string
}

// NewBase64Image creates a Base64Image from a base64 string
func NewBase64Image(base64Str []byte, name string) (*Base64Image, error) {

	// Get image size in bytes
	imageSize := len(base64Str)

	// Detect image type
	contentType := http.DetectContentType(base64Str)
	ext := ""
	switch contentType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	default:
		return nil, fmt.Errorf("unsupported image format")
	}

	// Decode the image to get dimensions
	img, _, err := image.Decode(bytes.NewReader(base64Str))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	return &Base64Image{
		Data: img,
		Size: imageSize,
		Ext:  ext,
		Name: name,
	}, nil
}

// GetImageSize returns the width, height, and size in bytes
func (b *Base64Image) GetImageSize() (int, int, int) {
	return b.Data.Bounds().Dx(), b.Data.Bounds().Dy(), b.Size
}

// GetImageExt returns the image extension
func (b *Base64Image) GetImageExt() string {
	return b.Ext
}

// GetImageName returns the image name
func (b *Base64Image) GetImageName(includeExt bool) string {
	if includeExt {
		return b.Name + b.Ext
	}
	return b.Name
}

// 判断图片是否为等比例图片
func (b *Base64Image) IsAspectRatioEqual() bool {
	return b.Data.Bounds().Dx() == b.Data.Bounds().Dy()
}

// ResizeImage resizes the image to the specified dimensions
func (b *Base64Image) ResizeImage(width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), b.Data, resize.Lanczos3)
}

func (b *Base64Image) ResizeImageAndSave(width, height int) (string, error) {
	resizedImage := resize.Resize(uint(width), uint(height), b.Data, resize.Lanczos3)

	ext := b.GetImageExt()
	newName := fmt.Sprintf("%s_%d_%d%s", b.Name, width, height, ext)

	file, err := os.Create(newName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, resizedImage, nil)
	case ".png":
		err = png.Encode(file, resizedImage)
	case ".gif":
		err = gif.Encode(file, resizedImage, &gif.Options{})
	default:
		return "", fmt.Errorf("unsupported file format: %s", ext)
	}

	if err != nil {
		return "", err
	}

	return newName, nil
}
