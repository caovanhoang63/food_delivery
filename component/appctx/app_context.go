package appctx

import (
	"food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider}
}

func (appContext *appContext) GetMainDbConnection() *gorm.DB {
	return appContext.db
}

func (appContext *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appContext.uploadProvider
}
