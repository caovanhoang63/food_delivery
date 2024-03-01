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

// v1/categories/

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := categorybiz.NewCreateCategoryBiz(store)
		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
