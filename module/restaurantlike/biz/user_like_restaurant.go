package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"log"
)

type LikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikeCountRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type likeRestaurantBiz struct {
	store    LikeRestaurantStore
	incStore IncLikeCountRestaurantStore
}

func NewLikeRestaurantBiz(store LikeRestaurantStore, incStore IncLikeCountRestaurantStore) *likeRestaurantBiz {
	return &likeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *likeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
