package main

import (
	"encoding/binary"
	"encoding/json"
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
	Name        string `form:"name" binding:"required"`
	Type        string `form:"type" binding:"required"`
	Industry    string `form:"industry" binding:"required"`
	Description string `form:"description" binding:"required"`
	City        string `form:"city" binding:"required"`
	Email       string `form:"email" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

func showbusregpage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "Регистрация бизнеса",
	})
}

func regbusiness(c *gin.Context) {
	var form BusinessForm
	bsn := Business{}
	if c.Bind(&form) == nil {
		db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Batch(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("business"))
			id, _ := b.NextSequence()
			bsn.Id = int(id)
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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Успех. Верификация",
		})
		// } else {
		// 	c.HTML(http.StatusOK, "index.html", gin.H{
		// 		"title": "Неуспех",
		// 	})
		// }
	}
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
