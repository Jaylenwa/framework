package driver

import (
	"context"
	"framework/interfaces/handler/dto"
)

// go:generate mockgen -source=./svc_user.go -destination ./mock/svc_user.go -package mock
type UserService interface {
	FindUserById(ctx context.Context, id int64) (res dto.GetUserByIdRsp, err error)
	FindUserList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []dto.GetUserListRsp, err error)
	CreateUser(ctx context.Context, req dto.CreateUserReq) (id int64, err error)
	UpdateUser(ctx context.Context, id int64, req dto.UpdateUserReq) (err error)
	DelUser(ctx context.Context, id int64) (err error)
}
