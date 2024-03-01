package gincategory

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	categorybiz "food-delivery/module/category/biz"
	categorystorage "food-delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSqlStore(appCtx.GetMainDbConnection())
		biz := categorybiz.NewDeleteCategoryBiz(store)
		if err := biz.DeleteCategory(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
