package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
