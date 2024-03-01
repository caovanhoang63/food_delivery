package categorystorage

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCategoryWithConditions(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*categorymodel.Category, error) {
	var data categorymodel.Category
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &data, nil
}
