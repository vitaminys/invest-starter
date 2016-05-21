package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var siteName string = "InvestStarter"

func main() {

	router := gin.Default()

	router.Static("css/", "./files/site/css/")
	router.Static("js/", "./files/site/js/")
	router.Static("img/", "./files/site/img/")
	router.LoadHTMLGlob("files/site/*.html")

	router.GET("/", mpage)
	router.GET("/business-reg", showbusregpage)
	router.POST("/business-reg/check", regbusiness)
	router.GET("/business-profile", busprofile)

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
		"name":  siteName,
	})
}
