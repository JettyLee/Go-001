package model

import "time"

type BaseModel struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
	Status  string `json:"status"`
	//创建人
	Creator    string    `json:"creator"`
	CreateTime time.Time `json:"create_time"`
	//修改人
	Editor     string    `json:"editor"`
	UpdateTime time.Time `json:"update_time"`
	//删除 标示 0正常  1已删除
	IsDelete   int       `json:"is_delete"`
	DeleteTime time.Time `json:"delete_time"`
	Remark     string    `json:"remark"`
}
