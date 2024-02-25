package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		err := c.ShouldBind(&data)
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDbConnection()

		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		err = biz.Create(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
