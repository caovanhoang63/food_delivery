package categorystorage

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db
	if err := db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
