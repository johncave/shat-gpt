package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")
	router.StaticFile("/favicon.ico", "favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:8080")
}
