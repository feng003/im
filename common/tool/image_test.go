package tool

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// 获取图片宽高和文件大小
func getImageSize(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, 0, err
	}

	const maxSize = 5 * 1024 * 1024 // 5MB
	if fileInfo.Size() > maxSize {
		return 0, 0, errors.New("file size exceeds 5MB")
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

// 判断图片是否为等比例图片
func isAspectRatioEqual(width, height int) bool {
	return width == height
}

// 压缩等比例图片到指定比例
func resizeImage(filePath string, newWidth, newHeight int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * img.Bounds().Dx() / newWidth
			srcY := y * img.Bounds().Dy() / newHeight
			resizedImg.Set(x, y, img.At(srcX, srcY))
		}
	}

	// 创建新文件名
	ext := filepath.Ext(filePath)
	newFileName := fmt.Sprintf("%s_%d%s", filePath[:len(filePath)-len(ext)], newWidth, ext)

	newFile, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	err = jpeg.Encode(newFile, resizedImg, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Resized image saved as: %s\n", newFileName)
	return nil
}

// 判断图片格式
func getImageFormat(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "",
			err
	}
	defer file.Close()

	_, format, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	return format, nil
}

func TestImage(t *testing.T) {
	filePath := "/home/feng003/yusheng/statics/03.png"

	width, height, err := getImageSize(filePath)
	if err != nil {
		fmt.Println("Error getting image size:", err)
		return
	}
	fmt.Printf("Image size: %dx%d\n", width, height)

	if isAspectRatioEqual(width, height) {
		fmt.Println("The image is a square.")
		err := resizeImage(filePath, 750, 750)
		if err != nil {
			fmt.Println("Error resizing image:", err)
			return
		}
	} else {
		fmt.Println("The image is not a square.")
	}

	format, err := getImageFormat(filePath)
	if err != nil {
		fmt.Println("Error getting image format:", err)
		return
	}
	fmt.Printf("Image format: %s\n", format)
}
