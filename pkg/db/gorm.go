package db

import (
	"go.uber.org/zap"
	pdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tarunngusain08/culturehub/pkg/db/postgres"
	"github.com/tarunngusain08/culturehub/pkg/log"
)

var logger = log.New("internal/db")

// NewDB creates a new postgres connection and returns gorm.DB
func NewDB(dns string, opts ...Option) (*gorm.DB, error) {
	logger.Info("connecting to database", zap.String("dns", dns))
	conf := new(gorm.Config)
	for _, o := range opts {
		o(conf)
	}

	db, err := gorm.Open(pdriver.Open(dns), conf)
	if err != nil {
		logger.Info("gorm connection error", zap.Error(err))
		return nil, err
	}
	return db, nil
}

func ResetDB() error {
	statements := []string{
		// "Create Type round_type AS ENUM ()",
		// "Create Type action AS ENUM ()",
	}
	err := postgres.ResetDB(statements)
	if err != nil {
		return err
	}
	return nil
}
