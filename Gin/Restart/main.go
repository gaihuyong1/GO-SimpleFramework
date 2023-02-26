package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/",func(c *gin.Context){
		time.Sleep(10*time.Second)
		c.String(http.StatusOK,"success")
	})

	ser:=&http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func(){
		if err:=ser.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
			log.Fatal("listen: %s\n",err)
		}
	}()

	quit:=make(chan os.Signal)
  signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("shutdown server")
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if err:=ser.Shutdown(ctx);err!=nil{
		log.Fatal("server shutdown: ",err)
	}
	log.Println("server exiting")
}