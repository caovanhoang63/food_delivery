package main

import (
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.FindRestaurantByID(appCtx))
	restaurants.GET("/", ginrestaurant.ListRestaurantWithCondition(appCtx))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
}
