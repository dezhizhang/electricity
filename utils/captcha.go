package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"log"
)

var store = base64Captcha.DefaultMemStore

func MakeCaptcha() (id string,base64 string,err error) {
	var driver base64Captcha.Driver

	driverString := base64Captcha.DriverString{
		Height: 40,
		Width: 120,
		NoiseCount: 0,
		ShowLineOptions: 2 | 4,
		Length: 4,
		Source: "13467ertyiadfhjkxcvbnERTYADFGHJKXCVBN",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver,store)
	id, base64, err = c.Generate()
	if err != nil {
		log.Fatalf("生成失败%v",err)
		return "", "", err
	}
	return id,base64,err
}