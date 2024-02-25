package restaurantbiz

import (
	"context"
	"food-delivery/common"
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
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
