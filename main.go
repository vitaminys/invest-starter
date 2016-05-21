package main

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
	router.POST("/business-reg/check/", regbusiness)
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
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))
		i := 0
		// l := 10
		arr := make([]*Business, 0)
		b.ForEach(func(k, v []byte) error {
			arr = append(arr, &Business{})
			// log.Printf("key=%s, value=%s\n", k, v)
			err = json.Unmarshal(v, arr[i])
			if (len(arr[i].Name) > 0) {
				i++
			}
			if err != nil {
				log.Println("JSON err:", err)
			}
			return nil
		})
		arr2 := arr[0:i]
		log.Println(arr2)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Главная страница",
			"name":       siteName,
			"Businesses": arr2,
		})
		return nil
	})
}
