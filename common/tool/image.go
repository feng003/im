package tool

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"  // 必须导入此包以启用 GIF 解码
	_ "image/jpeg" // 必须导入此包以启用 JPEG 解码
	_ "image/png"  // 必须导入此包以启用 PNG 解码
	"path/filepath"
	"strings"
)

// ImageTool 图片工具
type ImageTool struct {
}

// NewImageTool 获取图片工具实例
func NewImageTool() *ImageTool {
	return &ImageTool{}
}

// 下载图片并获取图片的尺寸
func (t *ImageTool) DownloadImageAndGetSize(imageUrl string) (width, height int, err error) {
	//fmt.Println("imageUrl = ", imageUrl)
	imageData, err := DownloadImageWithBytes(imageUrl)
	if err != nil {
		return
	}
	//fmt.Println("imageLen : ", len(imageData))
	return t.GetImageSize(imageData)
}

// GetImageSize 获取图片的尺寸
func (t *ImageTool) GetImageSize(imageData []byte) (width, height int, err error) {

	// 解析图片数据
	reader := bytes.NewReader(imageData)
	img, format, err := image.Decode(reader)
	if err != nil {
		return 0, 0, fmt.Errorf("error decoding image: %v", err)
	}

	// 获取图片尺寸
	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()

	// 确认图片格式
	if !strings.Contains("jpeg jpg png", format) {
		return 0, 0, fmt.Errorf("不支持的图片格式! image format: %v", format)
	}

	return width, height, nil
}

func ParseImageURL(url string, includeExt bool) string {
	filename := filepath.Base(url)

	if includeExt {
		return filename
	}

	// 不包含扩展名时 stripping extension
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
