package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()

	router.GET("/getJson",func(c *gin.Context){
		c.JSON(200,gin.H{
			"html":"<b>hello world</b>",
		})
	})

	router.GET("/getPurtJson",func(c *gin.Context){
		c.PureJSON(200,gin.H{
			"html":"<b>hello world</b>",
		})
	})

	router.Run(":8080")
}