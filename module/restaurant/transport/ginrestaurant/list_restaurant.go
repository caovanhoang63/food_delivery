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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.FullFill()

		err = c.ShouldBind(&filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := appCtx.GetMainDbConnection()

		store := restaurantstorage.NewSqlStore(db)

		biz := restaurantbiz.NewListRestaurantBiz(store)

		data, err = biz.List(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}

}
