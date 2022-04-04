package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type  BaseController struct {

}

func (that BaseController) success(c *gin.Context,message string,redirect string)  {
	c.HTML(http.StatusOK,"backend/public/success.html",gin.H{
		"message":message,
		"redirect":redirect,
	})
}

func (that BaseController) error(c *gin.Context,message string,redirect string) {
	c.HTML(http.StatusOK,"backend/public/success.html",gin.H{
		"message":message,
		"redirect":redirect,
	})
}