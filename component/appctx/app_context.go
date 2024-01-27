package appctx

import "gorm.io/gorm"

type appContext struct {
	db *gorm.DB
}

type AppContext interface {
	GetMainDbConnection() *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (appContext *appContext) GetMainDbConnection() *gorm.DB {
	return appContext.db
}
