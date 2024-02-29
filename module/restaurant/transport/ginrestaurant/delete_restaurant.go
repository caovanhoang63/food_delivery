package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDbConnection()
		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantStore(store, requester)

		if err = biz.Delete(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
