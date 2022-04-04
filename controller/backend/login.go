package backend

import (
	"electricity/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginController struct {

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

}
