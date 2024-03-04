package subcriber

import (
	"context"
	"food-delivery/component/appctx"
)

func SetUpPubSubSubcriber(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikeCountWhenUserLikeRestaurant(appCtx, ctx)
	DecreaseLikeCountWhenUserLikeRestaurant(appCtx, ctx)
}
