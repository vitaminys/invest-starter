package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware

	router := gin.Default()
	router.LoadHTMLGlob("files/*")

	router.GET("/", mpage)

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
		"title": "Main website",
	})
}

func re
