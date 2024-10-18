package models

import (
	"fmt"

	"gorm.io/gorm"
)

func (dao defaultDAO) Migrate() error {
	if dao.db == nil {
		return fmt.Errorf("db is nil")
	}
	err := dao.db.AutoMigrate(
		new(User),
		new(Idea),
		new(Activity),
	)
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
	Idea() IdeaDao
	Activity() ActivityDao
	User() UserDao
}

type (
	defaultDAO struct{ baseDAO }
	baseDAO    struct{ db *gorm.DB }
)

func (d *defaultDAO) Idea() IdeaDao         { return IdeaDao{baseDAO: d.baseDAO} }
func (d *defaultDAO) Activity() ActivityDao { return ActivityDao{baseDAO: d.baseDAO} }
func (d *defaultDAO) User() UserDao         { return UserDao{baseDAO: d.baseDAO} }
