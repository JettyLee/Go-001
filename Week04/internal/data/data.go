package data

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Data struct {
	gormDB  *gorm.DB
	redisDB *redis.Client
}

func ProvideData(gormDB *gorm.DB, redisDB *redis.Client) *Data {
	service := &Data{
		gormDB:  gormDB,
		redisDB: redisDB,
	}
	return service
}
