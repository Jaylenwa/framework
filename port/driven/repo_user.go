package portDriven

import (
	"context"
	"framework/infra/po"
	"framework/infra/utils/query"
)

// go:generate mockgen -source=./repo_user.go -destination ./mock/repo_user.go -package mock
type UserRepo interface {
	FindById(ctx context.Context, id int64) (res po.User, err error)
	FindByQuery(ctx context.Context, queries []*query.Query) (res po.User, err error)
	FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []po.User, err error)
	Insert(ctx context.Context, res po.User) (id int64, err error)
	Update(ctx context.Context, id int64, res po.User) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
