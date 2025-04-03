package imageTool

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type LocalImage struct {
	Path string
	Data image.Image
	Size int64
}

// NewImageTool 获取图片工具实例
func NewLocalImage(imagePath string) (*LocalImage, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &LocalImage{
		Path: imagePath,
		Data: img,
		Size: info.Size(),
	}, nil
}

// GetImageSize 获取图片的宽高
func (li *LocalImage) GetImageSize() (int, int, int64) {
	return li.Data.Bounds().Dx(), li.Data.Bounds().Dy(), li.Size
}

// GetImageExt 获取图片的扩展名
func (li *LocalImage) GetImageExt() string {
	return filepath.Ext(li.Path)
}

// GetImageName 获取图片的名字
func (li *LocalImage) GetImageName(includeExt bool) string {
	if includeExt {
		return filepath.Base(li.Path)
	}
	return filepath.Base(li.Path[:len(li.Path)-len(filepath.Ext(li.Path))])
}

// ResizeImage 对图片进行等比例压缩
func (li *LocalImage) ResizeImage(width, height int) image.Image {
	return resize.Resize(uint(width), uint(height), li.Data, resize.Lanczos3)
}

// ResizeImageAndSave 生成新的等比例压缩图片并保存
func (ri *LocalImage) ResizeImageAndSave(path string, width, height int) (string, error) {
	resizedImage := resize.Resize(uint(width), uint(height), ri.Data, resize.Lanczos3)

	ext := ri.GetImageExt()
	nameWithoutExt := ri.GetImageName(false)
	newName := fmt.Sprintf("%s_%d_%d%s", nameWithoutExt, width, height, ext)

	file, err := os.Create(path + "/" + newName)
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

// 判断图片是否为等比例图片
func IsAspectRatioEqual(width, height int) bool {
	return width == height
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
