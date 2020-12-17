package biz

import (
	"JettyPayGo/internal/data"
)

type Biz struct {
	data *data.Data
}

func ProvideBiz(data *data.Data) *Biz {
	return &Biz{data: data}
}
