package gincategory

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	categorybiz "food-delivery/module/category/biz"
	categorystorage "food-delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindCategoryById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := categorybiz.NewFindCategoryBiz(store)

		result, err := biz.FindById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
