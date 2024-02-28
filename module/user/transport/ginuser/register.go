package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/Hasher"
	"food-delivery/component/appctx"
	userbiz "food-delivery/module/user/biz"
	usermodel "food-delivery/module/user/model"
	userstore "food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
)

func RegisterUser(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appctx.GetMainDbConnection()

		store := userstore.NewSqlStore(db)

		md5Hasher := Hasher.NewMd5Hash()

		biz := userbiz.NewRegisterBiz(store, md5Hasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
