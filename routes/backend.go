package routes

import (
	"electricity/controller/backend"
	"github.com/gin-gonic/gin"
)

func BackendRouter(r *gin.Engine) {
	//backend := r.Group("/admin",middleware)
	b := r.Group("/admin")
	b.GET("/login",backend.LoginController{}.Login)
	b.GET("/captcha",backend.LoginController{}.Captcha)
	b.POST("/doLogin",backend.LoginController{}.DoLogin)

	b.GET("/role",backend.RoleController{}.Role)
	b.GET("/role/add",backend.RoleController{}.Add)
	b.POST("/role/doAdd",backend.RoleController{}.DoAdd)
}