package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {

}

func (that LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK,"backend/login/login.html",nil)
}
