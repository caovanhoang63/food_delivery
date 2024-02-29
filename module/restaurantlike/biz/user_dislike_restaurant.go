package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"log"
)

type DislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type dislikeRestaurantBiz struct {
	store    DislikeRestaurantStore
	decStore DecreaseLikeRestaurantStore
}

func NewDislikeRestaurantBiz(store DislikeRestaurantStore, decStore DecreaseLikeRestaurantStore) *dislikeRestaurantBiz {
	return &dislikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz *dislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, UserId, RestaurantId int) error {
	if err := biz.store.Delete(ctx, UserId, RestaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.decStore.DecreaseLikeCount(ctx, RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
