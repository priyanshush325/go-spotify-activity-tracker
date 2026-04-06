package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"spotify-activity-tracker/utils"
)

func main () { 
	router := gin.Default()
	router.GET("/ping", func (c *gin.Context) {
		fmt.Println("request received: %q", c.Request.Body)
		c.JSON(200, gin.H {
			"body": "pong",
		})
	})

	router.GET("/last-listened", func(c *gin.Context) {
		track, err := utils.GetLastListened()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, track)
	})

	router.Run()
}
