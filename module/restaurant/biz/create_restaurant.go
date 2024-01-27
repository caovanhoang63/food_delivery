package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantStoreBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantStoreBiz {
	return &createRestaurantStoreBiz{store: store}
}

func (biz *createRestaurantStoreBiz) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}

	err := biz.store.Create(context, data)
	if err != nil {
		return err
	}

	return nil
}
