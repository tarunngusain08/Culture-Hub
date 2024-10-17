package postgres

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"go.uber.org/zap"
)

func ResetDB(statements []string) error {
	db, err := sql.Open(driver(), ConnectionStr())
	if err != nil {
		logger.Error("db connection error", zap.Error(err))
		return err
	}
	err = dropDB(db)
	if err != nil {
		return err
	}
	err = createDB(db)
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		logger.Error("db close error", zap.Error(err))
		return err
	}

	db, err = sql.Open(driver(), ConnectionStrWithDB())
	if err != nil {
		logger.Error("db connection error", zap.Error(err))
		return err
	}

	err = createTypes(db, statements)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		logger.Error("db close error", zap.Error(err))
	}
	return err
}

func dropDB(db *sql.DB) error {
	logger.Info("dropping database", zap.String("dbname", DBName()))
	_, err := db.Exec("DROP DATABASE IF EXISTS " + DBName())
	if err != nil {
		logger.Error("drop db error", zap.Error(err))
	}
	return err
}

func createDB(db *sql.DB) error {
	logger.Info("creating database", zap.String("dbname", DBName()))
	_, err := db.Exec("CREATE DATABASE " + DBName())
	time.Sleep(2 * time.Second)
	if err != nil {
		logger.Error("create db error", zap.Error(err))
	}
	return err
}

func createTypes(db *sql.DB, statements []string) error {
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			if driverErr, ok := err.(*pq.Error); ok && driverErr.Code == "42710" {
				logger.Info("db create type: object already exists")
				return nil
			}
			logger.Error("error creating type", zap.Error(err))
			return err
		}
	}
	return nil
}
