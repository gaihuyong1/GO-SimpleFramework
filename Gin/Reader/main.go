package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/get",func(c *gin.Context){
		response,err:=http.Get("https://www.github.com")
		if err!=nil||response.StatusCode!=http.StatusOK{
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader:=response.Body
		contentLength:=response.ContentLength
		contentType:=response.Header.Get("Content-Type")
		extHeader:=map[string]string{
			"Disposition":"dispotion",
		}
		c.DataFromReader(http.StatusOK,contentLength,contentType,reader,extHeader)
	})

	router.Run(":8080")
}