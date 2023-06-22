package repositoryimpl

import (
	"context"
	"framework/global"
	"framework/infra/repository/po"
	"framework/infra/utils/query"
	"framework/port/driven"
	"gorm.io/gorm"
	"sync"
)

var (
	userRepoOnce sync.Once
	userRepoImpl driven.UserRepo
)

type userRepo struct {
	db *gorm.DB
}

var _ driven.UserRepo = &userRepo{}

func NewUserRepo() driven.UserRepo {
	userRepoOnce.Do(func() {
		userRepoImpl = &userRepo{
			db: global.GDB,
		}
	})
	return userRepoImpl
}

func (repo *userRepo) FindById(ctx context.Context, id int64) (user po.User, err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Where("id = ?", id).First(&user).Error
	return
}

func (repo *userRepo) FindByQuery(ctx context.Context, queries []*query.Query) (user po.User, err error) {
	tx := repo.db.WithContext(ctx)

	condition := query.GenerateQueryCondition(queries)

	err = tx.Where(condition).First(&user).Error
	return
}

func (repo *userRepo) FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []po.User, err error) {
	tx := repo.db.WithContext(ctx)

	limit := 10
	offset := 0

	condition := make(map[string]interface{})

	for k, v := range filter {
		if k == "limit" {
			limit = int(v.(float64))
		} else if k == "offset" {
			offset = int(v.(float64))
		} else {
			condition[k] = v
		}
	}

	dbQuery := tx.Model(&po.User{}).Where(condition)

	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	dbQuery = dbQuery.Count(&total)

	err = dbQuery.Limit(limit).Offset(offset).Find(&res).Error
	return
}

func (repo *userRepo) Insert(ctx context.Context, user po.User) (id int64, err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Create(&user).Error
	if err != nil {
		return
	}

	id = user.Id
	return
}

func (repo *userRepo) Update(ctx context.Context, id int64, user po.User) (err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Model(&po.User{}).Where("id = ?", id).Updates(&user).Error
	return
}

func (repo *userRepo) Delete(ctx context.Context, id int64) (err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Where("id = ?", id).Delete(&po.User{}).Error
	return
}
