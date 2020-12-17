package biz

import (
	"context"
)

func (b *Biz) GetUserPassword(ctx context.Context, name string) (string, error) {
	return b.data.GetUserPassword(ctx, name)
	//if err !=nil {
	//	err=b.data.SetUserPassword(ctx,name,"default password")
	//	if err!=nil{
	//		log.Println("设置默认密码失败：%v",err)
	//	}
	//}
	// reuslt,err
}
