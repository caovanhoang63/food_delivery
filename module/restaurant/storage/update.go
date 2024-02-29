package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *restaurantmodel.RestaurantUpdate) error {

	if err := s.db.Where("id = ? ", id).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
