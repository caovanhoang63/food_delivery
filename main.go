package main

import (
	"food-delivery/component/appctx"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//dsn aka connection string
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(db, err)
	}

	r := gin.Default() //create a server

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := appctx.NewAppContext(db)

	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.FindRestaurantByID(appCtx))
	restaurants.GET("/", ginrestaurant.ListRestaurantWithCondition(appCtx))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

	err = r.Run() //listen and serve

	if err != nil {
		log.Fatal(err)
	}

}
