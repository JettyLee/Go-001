package service

import (
	code "JettyPayGo/pkg/error"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func (s *Service) GetUserPassword(ctx *gin.Context) {

	result, err := s.biz.GetUserPassword(ctx, "abc")
	if err != nil && errors.Is(err, code.RedisNotFound) {
		log.Println("报错了:", errors.Unwrap(err))
		ctx.JSON(200, gin.H{
			"message": "没有查询到结果",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": result,
	})
	return
}
