package tool

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path"
	"path/filepath"
	"im/common/tool/imageTool"
	"im/common/uniqueid"
	"im/common/xerr"
)

func SaveFile(data []byte, opt ...map[string]string) (string, error) {
	// Generate a unique filename
	var prefix uniqueid.SnPrefix
	if len(opt[0]["prefix"]) > 0 {
		prefix = uniqueid.SnPrefix(opt[0]["prefix"])
	} else {
		prefix = uniqueid.SN_PREFIX_UPLOAD_PICTURE
	}

	fileType := GetFileExtension(data[:10])
	if len(fileType) <= 0 {
		return "", fmt.Errorf("failed to get file extension")
	}

	filename := uniqueid.GenSn(prefix) + "." + fileType

	domain := opt[0]["domain"]
	url := "/statics/" + opt[0]["project"] + "/" + opt[0]["module"]
	dir := path.Join(opt[0]["path"], url)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Save the base64-decoded image to a file
	err := os.WriteFile(filepath.Join(dir, filename), data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save base64 image to file: %v", err)
	}

	// Return the URL of the saved file
	return fmt.Sprintf("%s%s/%s", domain, url, filename), nil
}

func SaveBase64Image(data []byte, opt ...map[string]string) (string, error) {
	// Generate a unique filename
	var prefix uniqueid.SnPrefix
	if len(opt[0]["prefix"]) > 0 {
		prefix = uniqueid.SnPrefix(opt[0]["prefix"])
	} else {
		prefix = uniqueid.SN_PREFIX_UPLOAD_PICTURE
	}

	fileType := GetFileExtension(data[:10])
	if len(fileType) <= 0 {
		return "", fmt.Errorf("failed to get file extension")
	}

	filename := uniqueid.GenSn(prefix) + "." + fileType
	// 设置一个dir目录为项目的static目录 如果不存在就执行创建
	domain := opt[0]["domain"]
	url := "/statics/" + opt[0]["project"] + "/" + opt[0]["module"]
	dir := path.Join(opt[0]["path"], url)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Save the base64-decoded image to a file
	err := os.WriteFile(filepath.Join(dir, filename), data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save base64 image to file: %v", err)
	}

	// Return the URL of the saved file
	return fmt.Sprintf("%s%s/%s", domain, url, filename), nil
}

func Base64ImageVerify(data []byte) error {
	base64Image, err := imageTool.NewBase64Image(data, "")
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "failed to create base64 image %+v", err)
	}
	w, h, size := base64Image.GetImageSize()
	logx.WithContext(context.Background()).Infof("w: %d, h: %d, size: %d", w, h, size)

	const maxSize = 5 * 1024 * 1024 // 5MB
	if size > maxSize {
		return errors.Wrapf(xerr.NewErrCode(xerr.PMS_COMMERCE_IMAGE_SIZE_ERROR), "image size is too large %d", size)
	}
	if w != h {
		return errors.Wrapf(xerr.NewErrCode(xerr.PMS_COMMERCE_IMAGE_NOT_RATIO_ERROR), "image width and height are not equal %d and %d", w, h)
	}

	return nil
}

func SaveBase64ImageCommerce(data []byte, opt ...map[string]string) (string, error) {
	var prefix uniqueid.SnPrefix
	if len(opt[0]["prefix"]) > 0 {
		prefix = uniqueid.SnPrefix(opt[0]["prefix"])
	} else {
		prefix = uniqueid.SN_PREFIX_UPLOAD_PICTURE
	}

	fileType := GetFileExtension(data[:10])
	if len(fileType) <= 0 {
		return "", fmt.Errorf("failed to get file extension")
	}

	filename := uniqueid.GenSn(prefix) + "." + fileType
	// 设置一个dir目录为项目的static目录 如果不存在就执行创建
	domain := opt[0]["domain"]
	url := "/statics/" + opt[0]["project"] + "/" + opt[0]["module"]
	dir := path.Join(opt[0]["path"], url)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Save the base64-decoded image to a file
	localFile := filepath.Join(dir, filename)
	err := os.WriteFile(localFile, data, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save base64 image to file: %v", err)
	}
	local, err := imageTool.NewLocalImage(localFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to create local image")
	}
	//logx.WithContext(context.Background()).Infof("dir: %+v and filename is %+v", dir, filename)
	filename, err = local.ResizeImageAndSave(dir, 750, 750)
	if err != nil {
		return "", errors.Wrap(err, "failed to resize image")
	}

	// Return the URL of the saved file
	return fmt.Sprintf("%s%s/%s", domain, url, filename), nil
}
