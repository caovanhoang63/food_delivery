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

func UpdateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err := biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
