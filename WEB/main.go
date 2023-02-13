package main

import (
	"net/http"

	"web"
)

func main() {
	r := web.Default()
	r.GET("/", func(c *web.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *web.Context) {
		names := []string{"web"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
