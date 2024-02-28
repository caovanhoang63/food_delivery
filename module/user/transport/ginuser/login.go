package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/Hasher"
	"food-delivery/component/appctx"
	"food-delivery/component/tokenprovider/jwt"
	userbiz "food-delivery/module/user/biz"
	usermodel "food-delivery/module/user/model"
	userstorage "food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userLogin usermodel.UserLogin
		err := c.ShouldBind(&userLogin)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDbConnection()
		tokenProvider := jwt.NewJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()
		store := userstorage.NewSqlStore(db)
		md5 := Hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(appCtx, store, tokenProvider, md5, 60*60*24*30)

		account, err := biz.Login(c.Request.Context(), &userLogin)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
