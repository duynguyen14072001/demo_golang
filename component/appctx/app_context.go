package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
