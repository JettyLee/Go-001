// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"JettyPayGo/internal/biz"
	"JettyPayGo/internal/data"
	"JettyPayGo/internal/service"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitApp(gormDB *gorm.DB, redisDB *redis.Client) *App {

	wire.Build(data.ProvideData, biz.ProvideBiz, service.ProvideService, ProvideApp)
	return &App{}
}
