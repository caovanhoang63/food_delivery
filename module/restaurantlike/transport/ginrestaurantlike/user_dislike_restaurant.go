package ginrestaurantlike

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	restaurantlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
)

// POST v1/restaurants/:id/dislike

func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDbConnection())
		decStore := restaurantstorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := restaurantlikebiz.NewDislikeRestaurantBiz(store, decStore)
		if err := biz.UserDislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
