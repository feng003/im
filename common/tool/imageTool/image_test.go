package imageTool

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalImage(t *testing.T) {
	imgPath := "/home/feng003/image/03.png"
	localImage, err := NewLocalImage(imgPath)
	assert.NoError(t, err)

	width, height, size := localImage.GetImageSize()
	assert.True(t, width > 0)
	assert.True(t, height > 0)
	fmt.Printf("size: %d\n", size)

	ext := localImage.GetImageExt()
	assert.Equal(t, ".png", ext)

	nameWithExt := localImage.GetImageName(true)
	assert.Equal(t, "03.png", nameWithExt)

	nameWithoutExt := localImage.GetImageName(false)
	assert.Equal(t, "03", nameWithoutExt)

	resizedImg := localImage.ResizeImage(1000, 1000)
	assert.Equal(t, 1000, resizedImg.Bounds().Dx())
	assert.Equal(t, 1000, resizedImg.Bounds().Dy())

	//save, err := localImage.ResizeImageAndSave(1000, 1000)
	//assert.Equal(t, 1000, resizedImg.Bounds().Dx())
	//assert.Equal(t, 1000, resizedImg.Bounds().Dy())
	//fmt.Printf("save: %s\n", save)
	//_, err = os.Stat(save)
	//assert.NoError(t, err)
}

// Remote Image Tests
func TestRemoteImage(t *testing.T) {
	imgURL := "http://front.shyusheng.com.cn/statics/store/commerce/Ys2024092623584381813079.png"
	remoteImage, err := NewRemoteImage(imgURL)
	assert.NoError(t, err)

	width, height, size := remoteImage.GetImageSize()
	assert.True(t, width > 0)
	assert.True(t, height > 0)
	fmt.Printf("width %d, height %d size: %d\n", width, height, size)

	ext := remoteImage.GetImageExt()
	assert.Equal(t, ".png", ext)

	nameWithExt := remoteImage.GetImageName(true)
	assert.Equal(t, "Ys2024092623584381813079.png", nameWithExt)

	nameWithoutExt := remoteImage.GetImageName(false)
	assert.Equal(t, "Ys2024092623584381813079", nameWithoutExt)

	resizedImg := remoteImage.ResizeImage(100, 100)
	assert.Equal(t, 100, resizedImg.Bounds().Dx())
	assert.Equal(t, 100, resizedImg.Bounds().Dy())

	save, err := remoteImage.ResizeImageAndSave(1000, 1000)
	assert.Equal(t, 100, resizedImg.Bounds().Dx())
	assert.Equal(t, 100, resizedImg.Bounds().Dy())
	fmt.Printf("save: %s\n", save)
	_, err = os.Stat(save)
	assert.NoError(t, err)
}

func TestBase64Image(t *testing.T) {
	// Sample base64 image string (replace with a valid one for real testing)

	// Test GetImageSize
	// width, height, size := base64Image.GetImageSize()
	// assert.True(t, width > 0)
	// assert.True(t, height > 0)
	// assert.True(t, size > 0)
	// fmt.Printf("width %d, height %d size: %d\n", width, height, size)

	// // Test GetImageExt
	// ext := base64Image.GetImageExt()
	// assert.Contains(t, []string{".jpg", ".jpeg", ".png"}, ext)

	// // Test GetImageName
	// nameWithExt := base64Image.GetImageName(true)
	// assert.Equal(t, imageName+ext, nameWithExt)

	// nameWithoutExt := base64Image.GetImageName(false)
	// assert.Equal(t, imageName, nameWithoutExt)

	// // Test ResizeImage
	// resizedImage := base64Image.ResizeImage(100, 100)
	// assert.Equal(t, 100, resizedImage.Bounds().Dx())
	// assert.Equal(t, 100, resizedImage.Bounds().Dy())
}
