package main

import (
	"github.com/gin-gonic/gin"
) 

func main(){
	router:=gin.Default()
	router.GET("/getJson",func(c *gin.Context){
		data:=map[string]interface{}{
			"name":"gin",
			"id":"<1>",
		}
		c.AsciiJSON(200,data)
	})

	router.Run(":8080")
}