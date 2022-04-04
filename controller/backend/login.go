package backend

import "github.com/gin-gonic/gin"

type LoginController struct {

}

func (that LoginController) Login(c *gin.Context) {
	c.HTML("backend/login/index.html",nil)
}
