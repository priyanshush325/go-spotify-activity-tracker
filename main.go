package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main () { 
	router := gin.Default()
	router.GET("/ping", func (c *gin.Context) {
		fmt.Println("request received: %q", c.Request.Body)
		c.JSON(200, gin.H {
			"body": "pong",
		})
	})

	router.Run()
}
