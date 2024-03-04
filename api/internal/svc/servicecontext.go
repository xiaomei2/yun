package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wechat-ocr/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config: c,
		UserDb: db,
	}
}
