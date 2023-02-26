package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()

	router.MaxMultipartMemory=8<<10
	router.POST("/files",func(c *gin.Context){
		form,_:=c.MultipartForm()
		files:=form.File["files"]

		for _,file:=range files{
			log.Print(file.Filename)

			c.SaveUploadedFile(file,"xx/xx")
		}
		c.String(http.StatusOK,fmt.Sprintf("%d  files uoloaded",len(files)))
	})

	router.Run(":8080")
}