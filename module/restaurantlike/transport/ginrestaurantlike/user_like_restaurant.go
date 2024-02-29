package ginrestaurantlike

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	restaurantlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
)

// POST v1/restaurants/:id/like

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDbConnection())
		incStore := restaurantstorage.NewSqlStore(appCtx.GetMainDbConnection())

		biz := restaurantlikebiz.NewLikeRestaurantBiz(store, incStore)
		if err = biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
