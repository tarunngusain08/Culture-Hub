package models

import (
	"fmt"

	"gorm.io/gorm"
)

func (dao defaultDAO) Migrate() error {
	if dao.db == nil {
		return fmt.Errorf("db is nil")
	}
	err := dao.db.AutoMigrate()
	return err
}

var defaultDao *defaultDAO

func DAO(db *gorm.DB) DaoService {
	if defaultDao == nil {
		defaultDao = &defaultDAO{baseDAO: baseDAO{db: db}}
	}
	return defaultDao
}

type DaoService interface {
	Migrate() error
}

type (
	defaultDAO struct{ baseDAO }
	baseDAO    struct{ db *gorm.DB }
)
