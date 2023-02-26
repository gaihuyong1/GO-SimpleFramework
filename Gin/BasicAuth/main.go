package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets=gin.H{
	"a":gin.H{"email":"abc@mail.com","phone":"123"},
	"b":gin.H{"email":"def@mail.com","phone":"456"},
	"c":gin.H{"email":"ghi@mail.com","phone":"789"},
}

func main(){
	router:=gin.Default()
	authorized:=router.GET("/admin",gin.BasicAuth(gin.Accounts{
		"name":"a",
		"password":"1234",
	}))
	authorized.GET("/secrets",func(c *gin.Context){
		user:=c.MustGet(gin.AuthUserKey).(string)
		if secret,ok:=secrets[user];ok{
			c.JSON(http.StatusOK,gin.H{"user":user,"secret":secret})
		}else{
			c.JSON(http.StatusOK,gin.H{"user":user,"secret":"not found"})
		}
	})

	router.Run(":8080")
}

