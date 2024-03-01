package categorystorage

import (
	"context"
	"food-delivery/common"
	categorymodel "food-delivery/module/category/model"
)

func (s *sqlStore) ListCategoryWithCondition(ctx context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	if paging.FakeCursor != "" {
		uid, err := common.FromBase58(paging.NextCursor)
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db.Where("id < ?", uid.GetLocalID())
	} else {
		db.Offset(paging.GetOffSet())
	}

	if err := db.Limit(paging.Limit).Find(&result).Order("id desc").Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
