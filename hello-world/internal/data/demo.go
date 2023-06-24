package data

import "github.com/go-kratos/kratos/v2/log"

type DemoRepo struct {
	data   *Data
	logger *log.Logger
}

func NewDemoRepo(data *Data, logger *log.Logger) *DemoRepo {
	return &DemoRepo{data: data, logger: logger}
}



