package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantStoreBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantStore(store DeleteRestaurantStore) *deleteRestaurantStoreBiz {
	return &deleteRestaurantStoreBiz{store: store}
}

func (biz *deleteRestaurantStoreBiz) Delete(ctx context.Context, id int) error {

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return err
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
