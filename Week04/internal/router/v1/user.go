package v1

import (
	"JettyPayGo/internal/service"
	"github.com/gin-gonic/gin"
)

func UserV1(r *gin.RouterGroup, srv *service.Service) {
	r.Use() //middleware.JWT()
	{

		r.GET("/user", srv.GetUserPassword)

	}
}
