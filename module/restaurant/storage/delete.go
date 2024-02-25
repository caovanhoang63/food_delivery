package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Delete(
	context context.Context,
	id int,
) error {

	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
