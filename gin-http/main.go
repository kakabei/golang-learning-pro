// geektutu.com
// main.go

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, xyecho!")
	})

	// // 匹配 /player1/xyecho
	r.GET("/player1/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello  %s", name)
	})

	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "kakaxi")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(200, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// Map 参数
	r.POST("/post2", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向 Redirect
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/gameplayer", defaultHandler)
		v1.GET("/bag", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/gameplayer", defaultHandler)
		v2.GET("/bag", defaultHandler)
	}

	// 上传文件
	// 单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	// 批量上传
	r.POST("/batchupload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// curl "http://127.0.0.1:8080/posts?id=9876&page=7" -X POST -d 'username=kakaxi&password=1234'
	// curl -g "http://127.0.0.1:8080/post2?ids[kakaxi]=10086&ids[meixi]=10002" -X POST -d 'names[aa]=AAA&names[bb]=BBB'
	// curl -i http://127.0.0.1:8080/redirect

	// curl http://127.0.0.1:8080/v1/gameplayer
	// curl http://127.0.0.1:8080/v2/gameplayer
	// curl http://127.0.0.1:8080/v1/bag
	// curl http://127.0.0.1:8080/v2/bag

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
