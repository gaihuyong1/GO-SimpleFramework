package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func(c *gin.Context){
		t:=time.Now()
		c.Set("example","123")
		c.Next()

		latency:=time.Since(t)
		log.Println(latency)

		status:=c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/get",func(c *gin.Context){
		example:=c.MustGet("example").(string)
		log.Println(example)
	})

	router.Run(":8080")
}