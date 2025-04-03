package imageTool

import (
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path"
	"strings"

	"fmt"
	"image"
)

type RemoteImage struct {
	URL  string
	Data image.Image
	Size int64
}

// NewImageTool 获取图片工具实例
func NewRemoteImage(imageURL string) (*RemoteImage, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: %s", resp.Status)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	// Get size from Content-Length header
	size := resp.ContentLength
	fmt.Printf("resp %+v", resp)

	return &RemoteImage{
		URL:  imageURL,
		Data: img,
		Size: size,
	}, nil
}

// GetImageSize 获取图片的宽高和大小
func (ri *RemoteImage) GetImageSize() (int, int, int64) {
	width := ri.Data.Bounds().Dx()
	height := ri.Data.Bounds().Dy()
	return width, height, ri.Size
}

// GetImageExt 获取图片的扩展名
func (ri *RemoteImage) GetImageExt() string {
	return path.Ext(ri.URL)
}

// GetImageName 获取图片的名字
func (ri *RemoteImage) GetImageName(includeExt bool) string {
	base := path.Base(ri.URL)
	if includeExt {
		return base
	}
	return strings.TrimSuffix(base, path.Ext(base))
}

// ResizeImage 对图片进行等比例压缩
func (li *RemoteImage) ResizeImage(width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), li.Data, resize.Lanczos3)
}

// ResizeImageAndSave 生成新的等比例压缩图片并保存
func (ri *RemoteImage) ResizeImageAndSave(width, height int) (string, error) {
	resizedImage := resize.Resize(uint(width), uint(height), ri.Data, resize.Lanczos3)

	ext := ri.GetImageExt()
	nameWithoutExt := ri.GetImageName(false)
	newName := fmt.Sprintf("%s_%d_%d%s", nameWithoutExt, width, height, ext)

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

// 下载图片并获取图片的尺寸
//func (t *ImageTool) DownloadImageAndGetSize(imageUrl string) (width, height int, err error) {
//	//fmt.Println("imageUrl = ", imageUrl)
//	imageData, err := DownloadImageWithBytes(imageUrl)
//	if err != nil {
//		return
//	}
//	//fmt.Println("imageLen : ", len(imageData))
//	return t.GetImageSize(imageData)
//}

// GetImageSize 获取图片的尺寸
//func (t *ImageTool) GetImageSize(imageData []byte) (width, height int, err error) {
//
//	// 解析图片数据
//	reader := bytes.NewReader(imageData)
//	img, format, err := imageTool.Decode(reader)
//	if err != nil {
//		return 0, 0, fmt.Errorf("error decoding imageTool: %v", err)
//	}
//
//	// 获取图片尺寸
//	bounds := img.Bounds()
//	width = bounds.Dx()
//	height = bounds.Dy()
//
//	// 确认图片格式
//	if !strings.Contains("jpeg jpg png", format) {
//		return 0, 0, fmt.Errorf("不支持的图片格式! imageTool format: %v", format)
//	}
//
//	return width, height, nil
//}
