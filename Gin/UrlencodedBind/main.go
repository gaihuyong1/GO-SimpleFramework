package main

import (
	"github.com/gin-gonic/gin"
)

type param struct{
	Name string `form:"name" binding:"required"`
	Id string `form:"id" binding:"required"`
}

func main(){
	router:=gin.Default()
	router.POST("/login",func(c *gin.Context){
		var pa param
		if c.ShouldBind(&pa)==nil{
			if pa.Name=="name"&&pa.Id=="id"{
				c.JSON(200,gin.H{"status":"success"})
			}else{
				c.JSON(401,gin.H{"status":"fail"})
			}
		}
	})

	router.Run(":8080")
}