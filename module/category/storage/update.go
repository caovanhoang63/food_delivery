package categorystorage

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {

	if err := s.db.Table(categorymodel.CategoryUpdate{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
