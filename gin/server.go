package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.GET("/data", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("Some binary data here."))
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"hello": "json"})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
