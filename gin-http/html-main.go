package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("html/*")
	//r.LoadHTMLGlob("html/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "火影已经完结了", "name": "我是卡卡西"})
	})
	r.Run()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
