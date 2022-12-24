package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("测试")
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		json := make(map[string]any) //注意该结构接受的内容
		err := c.BindJSON(&json)
		if err != nil {
			return
		}
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"text":     json["text"],
			"password": json["password"],
		})

	})
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err:", err)
	}
}
