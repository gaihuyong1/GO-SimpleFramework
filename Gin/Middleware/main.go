package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/middleware",MyMiddleware())

	authorized:=router.Group("/")
	authorized.Use(MyMiddleware())
	{
		authorized.POST("/login",func(c *gin.Context){})
		authorized.POST("/submit",func(c *gin.Context){})
	}

	router.Run(":8080")
}

func MyMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){}
}