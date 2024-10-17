package db

import "gorm.io/gorm"

type Option func(*gorm.Config)

func FKConstraintWhenMigrating(set bool) Option {
	return func(c *gorm.Config) {
		c.DisableForeignKeyConstraintWhenMigrating = !set
	}
}
