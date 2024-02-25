package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDbConnection()
		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantStore(store)

		if err = biz.Delete(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
