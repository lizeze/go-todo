package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {

	r := gin.Default()
	r.Static("/assets", "./web/assets")
	r.LoadHTMLFiles("web/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
