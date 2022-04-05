package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context) {
	// 1获取url访问地址
	pathname := strings.Split(c.Request.URL.String(),"?")[0]
	fmt.Println("pathname",pathname)
}