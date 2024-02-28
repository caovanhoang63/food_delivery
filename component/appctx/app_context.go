package appctx

import (
	"food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider, secretKey: secretKey}
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
