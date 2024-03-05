package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	"food-delivery/pubsub"
	"log"
)

type LikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikeCountRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type likeRestaurantBiz struct {
	store LikeRestaurantStore
	//incStore IncLikeCountRestaurantStore
	ps pubsub.Pubsub
}

func NewLikeRestaurantBiz(store LikeRestaurantStore, ps pubsub.Pubsub) *likeRestaurantBiz {
	return &likeRestaurantBiz{store: store, ps: ps}
}

func (biz *likeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicIncreaseLikeCountWhenUserLikeRestaurant,
		pubsub.NewMessage(data)); err != nil {
		log.Println("Err:", err)
	}

	return nil
}
