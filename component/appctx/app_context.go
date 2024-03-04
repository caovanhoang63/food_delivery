package appctx

import (
	"food-delivery/component/uploadprovider"
	"food-delivery/pubsub"
	"gorm.io/gorm"
)

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider, secretKey: secretKey, ps: ps}
}

func (appContext *appContext) GetMainDbConnection() *gorm.DB {
	return appContext.db
}

func (appContext *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appContext.uploadProvider
}

func (appContext *appContext) GetSecretKey() string {
	return appContext.secretKey
}

func (appContext *appContext) GetPubSub() pubsub.Pubsub { return appContext.ps }
