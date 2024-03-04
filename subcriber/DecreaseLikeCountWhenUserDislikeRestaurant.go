package subcriber

import (
	"context"
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"log"
)

func DecreaseLikeCountWhenUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicDecreaseLikeCountWhenUserDislikeRestaurant)

	go func() {
		common.AppRecover()
		for {
			msg := <-c
			id := msg.Data().(HasRestaurantId).GetRestaurantId()
			db := appCtx.GetMainDbConnection()
			if err := restaurantstorage.NewSqlStore(db).DecreaseLikeCount(ctx, id); err != nil {
				log.Println("err:", err)
			}
		}
	}()
}
