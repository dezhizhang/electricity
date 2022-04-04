package main

import (
	"electricity/routes"
	"github.com/gin-gonic/gin"
)

func main()  {

	r := gin.Default()

	//配置模板
	r.LoadHTMLGlob("templates/**/**/*")
	// 配置静态资源
	r.Static("/static","./static")

	routes.BackendRouter(r)



	r.Run(":8000")




}
