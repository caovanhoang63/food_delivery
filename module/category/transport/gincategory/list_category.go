package gincategory

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	categorybiz "food-delivery/module/category/biz"
	categorymodel "food-delivery/module/category/model"
	categorystorage "food-delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []categorymodel.Category
		var paging common.Paging
		var filter categorymodel.Filter

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := categorybiz.NewListCategoryBiz(store)
		data, err := biz.ListCategory(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))

	}
}
