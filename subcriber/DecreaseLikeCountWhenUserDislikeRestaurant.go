package subcriber

import (
	"context"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"food-delivery/pubsub"
)

func DecreaseLikeCountWhenUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Decrease like count when user dislikes restaurant",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDbConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
