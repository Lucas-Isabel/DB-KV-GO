package routes

import (
	"db-kv-go/server"

	"github.com/gin-gonic/gin"
)

func RunRoutes(s *server.Server) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong-due",
		})
	})
	r.POST("/set", s.SETk)
	r.GET("/get", s.GETk)
	r.GET("/all", s.ALLkv)
	r.DELETE("/delete", s.DELETEk)

	return r
}
