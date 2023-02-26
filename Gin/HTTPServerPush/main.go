package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main(){
	router:=gin.Default()
	router.GET("/",func(c *gin.Context){
		if pusher:=c.Writer.Pusher();pusher!=nil{
			if err:=pusher.Push("/assset/index.js",nil);err!=nil{
				log.Printf("faild to push")
			}
		}
		c.HTML(200,"https",gin.H{
			"statuc":"success",
		})
	})

	router.RunTLS(":8080","xxx/xx.pem","xxx/xxx.key")
}