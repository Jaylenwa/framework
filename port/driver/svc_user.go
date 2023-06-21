package driver

import (
	"context"
	"framework/interfaces/handler/dto"
)

// go:generate mockgen -source=./svc_user.go -destination ./mock/svc_user.go -package mock
type UserService interface {
	GetUserById(ctx context.Context, id uint64) (res dto.GetUserByIdRsp, err error)
	GetUserList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []dto.GetUserListRsp, err error)
	AddUser(ctx context.Context, req dto.AddUserReq) (id uint64, err error)
	UpdateUser(ctx context.Context, id uint64, req dto.UpdateUserReq) (err error)
	DelUser(ctx context.Context, id uint64) (err error)
}
