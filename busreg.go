package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	// "reflect"
	// "time"
)

type Business struct {
	Id          int
	Name        string
	Type        string
	Industry    string
	Description string
	City        string
	Email       string
	Password    string
	Verified    bool
}
type BusinessLight struct {
	Id          int
	Name        string
	Type        string
	Industry    string
	Description string
	City        string
	Verified    bool
}

type BusinessForm struct {
	Name        string `form:"name" binding:"required"`
	Type        string `form:"type" binding:"required"`
	Industry    string `form:"industry" binding:"required"`
	Description string `form:"description" binding:"required"`
	City        string `form:"city" binding:"required"`
	Email       string `form:"email" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

func showbusregpage(c *gin.Context) {
	c.HTML(http.StatusOK, "register_business.html", gin.H{
		"title": "Регистрация бизнеса",
		"name":  siteName,
	})
}

func regbusiness(c *gin.Context) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Batch(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		id, err := b.NextSequence()
		k := 0;
		if err != nil {
			log.Println(err)
			k = 0
		} else {
			k = int(id)
		}
		buf, err := json.Marshal(gin.H{
			"Id" : k,
			"Name" : c.PostForm("name"),
			"Type" : c.PostForm("type"),
			"Industry" : c.PostForm("industry"),
			"Description" : c.PostForm("descr"),
			"City": c.PostForm("city"),
			"Email": c.PostForm("email"),
			"Password": c.PostForm("passw"),
			"Verified": false,
		})
		if err != nil {
			return err
		}
		return b.Put(itob(k), buf)
	})
	if err != nil {
		log.Fatalln("Error in database while business registration")
	}
	c.Redirect(http.StatusMovedPermanently, "/business-profile")
}

func busprofile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile_business.html", gin.H{
		"title": "Найстройки профиля",
		"name":  siteName,
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
