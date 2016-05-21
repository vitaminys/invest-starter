package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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

type BusinessForm struct {
	Name        string //`form:"name" binding:"required"`
	Type        string //`form:"type" binding:"required"`
	Industry    string //`form:"industry" binding:"required"`
	Description string //`form:"description" binding:"required"`
	City        string //`form:"city" binding:"required"`
	Email       string //`form:"email" binding:"required"`
	Password    string //`form:"password" binding:"required"`
}

func showbusregpage(c *gin.Context) {
	c.HTML(http.StatusOK, "register_business.html", gin.H{
		"title": "Регистрация бизнеса",
		"name":  siteName,
	})
}

func regbusiness(c *gin.Context) {
	var form BusinessForm
	bsn := Business{}
	if c.Bind(&form) == nil {
		log.Println("register_business")
		db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Batch(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			// b := tx.Bucket([]byte("business"))
			id, err := b.NextSequence()
			if err != nil {
				log.Println(err)
				bsn.Id = 0
			} else {
				bsn.Id = int(id)
			}
			buf, err := json.Marshal(bsn)
			if err != nil {
				return err
			}
			return b.Put(itob(bsn.Id), buf)
		})

		if err != nil {
			log.Fatalln("Error in database while business registration")
		}

		// if form.User == "user" && form.Password == "password" {
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"title": "Успех. Верификация",
		// })
		// } else {
		// 	c.HTML(http.StatusOK, "index.html", gin.H{
		// 		"title": "Неуспех",
		// 	})
		// }
		c.Redirect(http.StatusMovedPermanently, "/business-profile")
	}
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
