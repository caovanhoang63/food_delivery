package subcriber

import (
	"context"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"food-delivery/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseLikeCountWhenUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Increase like count when user likes restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDbConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
