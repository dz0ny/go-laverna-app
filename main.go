package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	_static := http.FileServer(assetFS())
	r.Use(func(c *gin.Context) {
		_static.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	var port = ":3000"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
