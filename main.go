package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"log"
	"net/http"
)

func main() {
	fmt.Println("测试")
	router := gin.Default()

	var dbPath = "ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	// 获取 IP
	router.GET("/ip", func(c *gin.Context) {
		ip := c.ClientIP()
		c.String(http.StatusOK, ip)
	})

	// 获取 IP ip2region 数据
	router.GET("/ip2region", func(c *gin.Context) {
		ip := c.ClientIP()
		data, err := searcher.SearchByStr(ip)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, data)
	})

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
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("err:", err)
	}
}
