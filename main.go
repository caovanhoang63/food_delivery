package main

import (
	"context"
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/pubsub/pblocal"
	"food-delivery/subcriber"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	RegisterImageFormat()
	//dsn aka connection string
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(db, err)
	}
	db = db.Debug()

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3ApiKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := os.Getenv("S3_DOMAIN")
	systemSecretKey := os.Getenv("SYSTEM_SECRET_KEY")

	s3Provider := uploadprovider.NewS3Provider(
		s3BucketName,
		s3Region,
		s3ApiKey,
		s3Secret,
		s3Domain,
	)
	ps := pblocal.NewPubSub()

	appCtx := appctx.NewAppContext(db, s3Provider, systemSecretKey, ps)

	r := gin.Default() //create a server
	subcriber.SetUpPubSubSubcriber(appCtx, context.Background())

	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//setup route
	v1 := r.Group("/v1")
	setupRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)

	//listen and serve
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
