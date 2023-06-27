package entity

import (
	"context"
	"framework/infra/po"
	"framework/infra/utils/struct"
)

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
	Home string `json:"home"`
}

func (do User) ToPO(ctx context.Context) (po po.User, err error) {

	err = _struct.CopyStruct(&po, do)

	return
}
