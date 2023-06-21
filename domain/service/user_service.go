package service

import (
	"context"
	"framework/infra/repository/po"
	"framework/infra/repository/repositoryimpl"
	"framework/infra/utils/query"
	"framework/infra/utils/struct"
	"framework/interfaces/handler/dto"
	"framework/port/driven"
	"framework/port/driver"
	"sync"
)

var (
	userServiceOnce sync.Once
	userServiceImpl driver.UserService
)

type userService struct {
	userRepo driven.UserRepo
}

var _ driver.UserService = &userService{}

func NewUserService() driver.UserService {
	userServiceOnce.Do(func() {
		userServiceImpl = &userService{
			userRepo: repositoryimpl.NewUserRepo(),
		}
	})
	return userServiceImpl

}

func (svc *userService) GetUserById(ctx context.Context, id uint64) (res dto.GetUserByIdRsp, err error) {
	userPo, err := svc.userRepo.FindById(ctx, id)
	if err != nil {
		return
	}

	// PO_to_DO
	err = _struct.CopyStruct(&res, userPo)

	return
}

func (svc *userService) GetUserByQuery(ctx context.Context, queries []*query.Query) (res dto.GetUserByQueryRsp, err error) {
	userPo, err := svc.userRepo.FindByQuery(ctx, queries)

	// PO_to_DO
	err = _struct.CopyStruct(&res, userPo)

	return
}

func (svc *userService) GetUserList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []dto.GetUserListRsp, err error) {
	total, userList, err := svc.userRepo.FindList(ctx, filter, args...)

	// POs_to_DOs
	res = make([]dto.GetUserListRsp, 0)

	for _, user := range userList {
		do := dto.GetUserListRsp{}

		err = _struct.CopyStruct(&do, user)
		if err != nil {
			return
		}

		res = append(res, do)
	}

	return
}

func (svc *userService) AddUser(ctx context.Context, req dto.AddUserReq) (id uint64, err error) {

	var userPo po.User
	// DO_to_PO
	err = _struct.CopyStruct(&userPo, req)
	if err != nil {
		return
	}

	id, err = svc.userRepo.Insert(ctx, userPo)

	return
}

func (svc *userService) UpdateUser(ctx context.Context, id uint64, req dto.UpdateUserReq) (err error) {

	var userPo po.User
	// DO_to_PO
	err = _struct.CopyStruct(&userPo, req)
	if err != nil {
		return
	}

	err = svc.userRepo.Update(ctx, id, userPo)

	return
}

func (svc *userService) DelUser(ctx context.Context, id uint64) (err error) {
	err = svc.userRepo.Delete(ctx, id)

	return
}