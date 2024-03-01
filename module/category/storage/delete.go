package categorystorage

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
