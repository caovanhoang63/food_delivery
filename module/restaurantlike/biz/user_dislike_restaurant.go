package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"food-delivery/pubsub"
	"log"
)

type DislikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseLikeRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type dislikeRestaurantBiz struct {
	store DislikeRestaurantStore
	//decStore DecreaseLikeRestaurantStore
	ps pubsub.Pubsub
}

func NewDislikeRestaurantBiz(store DislikeRestaurantStore, ps pubsub.Pubsub) *dislikeRestaurantBiz {
	return &dislikeRestaurantBiz{store: store, ps: ps}
}

func (biz *dislikeRestaurantBiz) UserDislikeRestaurant(ctx context.Context, UserId, RestaurantId int) error {
	if err := biz.store.Delete(ctx, UserId, RestaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicDecreaseLikeCountWhenUserDislikeRestaurant,
		pubsub.NewMessage(&restaurantlikemodel.Like{RestaurantId: RestaurantId})); err != nil {
		log.Println("Err: ", err)
	}

	return nil
}
