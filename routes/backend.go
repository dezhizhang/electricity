package routes

import (
	"electricity/controller/backend"
	"github.com/gin-gonic/gin"
)

func BackendRouter(r *gin.Engine) {
	//backend := r.Group("/admin",middleware)
	admin := r.Group("/admin")
	admin.GET("/login",backend.LoginController{}.Login)
	admin.GET("/captcha",backend.LoginController{}.Captcha)
	admin.POST("/doLogin",backend.LoginController{}.DoLogin)
}