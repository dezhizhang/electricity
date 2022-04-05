package backend

import (
	"electricity/driver"
	"electricity/model"
	"electricity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (that LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK,"backend/login/login.html",nil)
}

func (that LoginController) Captcha(c *gin.Context)  {
	id,base64,err := utils.MakeCaptcha()
	if err != nil {
		log.Fatalf("生成验证码失败%v",err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"id":id,
		"image":base64,
	})
}

//登录

func (that LoginController) DoLogin(c *gin.Context)  {
	captchaId := c.PostForm("captchaId")
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")

	flag := utils.VerifyCaptcha(captchaId,code)
	if !flag {
		that.Error(c,"验证码不正确","/admin/login")
		return
	}

	fmt.Println("password",utils.Md5String(password))

	var user []model.User
	driver.DB.Where("email=? AND password=?",email,utils.Md5String(password)).Find(&user)

	if len(user) <=0 {
		that.Error(c,"用户名或密码错误","/admin/login")
		return
	}
	that.Success(c,"登录成功","/admin/manager")
}
