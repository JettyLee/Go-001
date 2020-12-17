package data

import (
	code "JettyPayGo/pkg/error"
	"context"
	xerrors "github.com/pkg/errors"
	"time"
)

func (d *Data) GetUserPassword(ctx context.Context, name string) (string, error) {
	reuslt, err := d.redisDB.Get(ctx, name).Result()
	if err != nil {
		return "", xerrors.Wrapf(code.RedisNotFound, "redis查询用户%s-密码报错:%v", name, err)
	}
	return reuslt, err
}

func (d *Data) SetUserPassword(ctx context.Context, name, password string) error {
	err := d.redisDB.Set(ctx, name, password, time.Duration(100*time.Second)).Err()
	if err != nil {
		return xerrors.Wrapf(err, "redis设置%s-用户迷药报错:%v", name, err)
	}
	return nil
}
