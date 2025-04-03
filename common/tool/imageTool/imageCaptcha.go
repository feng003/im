package imageTool

import "github.com/mojocn/base64Captcha"

// 图片验证码驱动
var imageCaptchaDriver *base64Captcha.DriverString

// CreateImageCaptcha 生成图片验证码
func CreateImageCaptcha() (imageId, imageBase64, imageCode string, err error) {
	// 判断有没有设置驱动
	if imageCaptchaDriver == nil {
		imageCaptchaDriver = &base64Captcha.DriverString{
			Height:          40,
			Width:           100,
			NoiseCount:      0,
			ShowLineOptions: 2,
			Length:          4,
			Source:          "1234567890",
		}
	}

	driver := imageCaptchaDriver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return c.Generate()
}

// VerifyImageCaptcha 校验图片验证码
func VerifyImageCaptcha(imageId, imageCode string) bool {
	return base64Captcha.DefaultMemStore.Verify(imageId, imageCode, true)
}
