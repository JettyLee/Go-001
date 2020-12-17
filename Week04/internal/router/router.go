package router

import (
	v1 "JettyPayGo/internal/router/v1"
	"JettyPayGo/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter(svc *service.Service) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userV1G := r.Group("/user/v1")
	v1.UserV1(userV1G, svc)
	return r
}
