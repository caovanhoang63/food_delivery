package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	err := biz.store.Update(ctx, id, data)
	if err != nil {
		return err
	}
	return nil
}
