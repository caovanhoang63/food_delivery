package userstorage

import (
	"context"
	"food-delivery/common"
	usermodel "food-delivery/module/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {

	db := s.db.Begin()

	if err := db.Table(usermodel.User{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}
	return nil
}
