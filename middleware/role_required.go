package middleware

import (
	"errors"
	"food-delivery/common"
	"food-delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func RoleRequired(ctx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for _, role := range allowRoles {
			if u.GetRole() == role {
				c.Set(common.CurrentUser, u)
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(errors.New("no permission")))
	}
}
