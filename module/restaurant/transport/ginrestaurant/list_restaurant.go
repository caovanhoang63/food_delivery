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

func ListRestaurantWithCondition(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data []restaurantmodel.Restaurant
		var paging common.Paging
		var filter restaurantmodel.Filter

		err := c.ShouldBind(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()
		err = c.ShouldBind(&filter)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDbConnection()
		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		data, err = biz.List(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}

}
