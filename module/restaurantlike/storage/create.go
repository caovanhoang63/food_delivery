package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
