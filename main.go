package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/contrib/static"
	"net/http"
	// "log"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware

	router := gin.Default()

	// router.StaticFS("files/site/", http.Dir("./"))
	router.Static("css/", "./files/site/css/")
	router.Static("js/", "./files/site/js/")
	router.Static("img/", "./files/site/img/")
	// router.Static("/img", "./files/site/img/")
	// router.LoadHTMLGlob("files/site/**/*")
	router.LoadHTMLGlob("files/site/*.html")
	// router.Use(static.Serve("/", static.LocalFile("/files/site", false)))
	// router.NotFound(static.Serve("/", "/files/site"))
	// router.Use(static.Serve("/", static.LocalFile("/files/site", false)))

	router.GET("/", mpage)
	router.GET("/business_reg", showbusregpage)
	router.POST("/business_reg/check", regbusiness)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}

// func mpage(c *gin.Context) {
// 	content := gin.H{"Hello": "World"}
// 	c.JSON(200, content)
// }

func mpage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Главная страница",
	})
}
